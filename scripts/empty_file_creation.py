import os

def create_files_from_list(source_folder, files, package_name):
    """
    Create new files based on the provided list of files in the source folder.
    
    Parameters:
    - source_folder: The folder where the new files will be created.
    - files: List of new files to be created.
    - package_name: Name of the package to be added to the files.
    """
    
    # For each file in the provided list, check if it exists. If not, create a new file with the package line
    created_count = 0
    for file in files:
        new_file_path = os.path.join(source_folder, file)
        
        if not os.path.exists(new_file_path):
            with open(new_file_path, 'w') as new_file:
                new_file.write(f"package {package_name}\n")
            created_count += 1
        else:
            print(f"File {file} already exists. Skipping.")

    print(f"Created {created_count} new files in {source_folder}")

if __name__ == "__main__":
    # Interactively ask user for inputs
    source_folder = input("Enter the path to the source folder where new files will be created: ")
    files_input = input("Enter the list of new files to be created (separated by space): ")
    files = files_input.split()
    package_name = input("Please enter the package name: ")
    
    create_files_from_list(source_folder, files, package_name)
