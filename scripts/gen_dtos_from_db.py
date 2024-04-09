import os
import psycopg2
from dotenv import load_dotenv

def make_singular(word):
    """Convert common plural words to singular."""
    if word.endswith("ies"):
        return word[:-3] + "y"
    elif word.endswith("s"):
        return word[:-1]
    return word

def parse_database_url(database_url):
    """Extract the database connection details from the DATABASE_URL string."""
    parts = database_url.split()
    details = {}
    for part in parts:
        key, value = part.split('=')
        details[key] = value
    return details

def connect_to_db(details):
    conn = psycopg2.connect(
        host=details["host"],
        port=details["port"],
        dbname=details["dbname"],
        user=details["user"],
        password=details["password"]
    )
    return conn

def to_pascal_case(snake_str, singularize=False):
    """Convert a snake_case string to PascalCase."""
    if singularize:
        snake_str = make_singular(snake_str)
    return ''.join(word.capitalize() for word in snake_str.split('_'))



def fetch_tables(cursor):
    query = "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
    cursor.execute(query)
    return cursor.fetchall()

def fetch_columns(cursor, table):
    query = "SELECT column_name, data_type FROM information_schema.columns WHERE table_name = %s;"
    cursor.execute(query, [table])
    return cursor.fetchall()

def generate_go_struct(table, columns):
    singular_table = to_pascal_case(table, True)
    imports = set()
    struct_str = "package models\n\n"
    for _, dtype in columns:
        if "timestamp" in dtype:
            imports.add("time")
    
    if "time" in imports:
        struct_str += "import (\n"
        for imp in imports:
            struct_str += f"\t\"{imp}\"\n"
        struct_str += ")\n\n"

    struct_str += f"type {singular_table}DTO struct {{\n"
    for column, dtype in columns:
        go_type = 'string'  # Default type
        if "integer" in dtype:
            go_type = 'int'
        elif "character varying" in dtype or "text" in dtype:
            go_type = 'string'
        elif "boolean" in dtype:
            go_type = 'bool'
        elif "timestamp" in dtype:
            go_type = 'time.Time'
        elif "ARRAY" in dtype:
            go_type = '[]string'
        struct_str += f"\t{to_pascal_case(column)} {go_type} `json:\"{column}\" db:\"{column}\"`\n"
    struct_str += "}\n"
    return struct_str


def write_to_file(output_dir, table, go_struct):
    singular_table = make_singular(table)
    file_name = f"{singular_table}_dto.go"
    file_path = os.path.join(output_dir, file_name)
    if not os.path.exists(file_path):  # Check if file already exists
        with open(file_path, "w") as f:
            f.write(go_struct)

if __name__ == "__main__":
    # Load environment variables from .env file
    load_dotenv()

    database_url = os.getenv("DATABASE_URL")
    details = parse_database_url(database_url)

    # Connect to the database
    conn = connect_to_db(details)

    # Fetch tables
    cursor = conn.cursor()
    tables = fetch_tables(cursor)

    # Ask user for destination folder
    output_dir = input("Enter the destination folder path: ")
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    # Generate Go structs for each table and write to separate files
    for table_tuple in tables:
        table = table_tuple[0]
        columns = fetch_columns(cursor, table)
        go_struct = generate_go_struct(table, columns)
        write_to_file(output_dir, table, go_struct)

    cursor.close()
    conn.close()
