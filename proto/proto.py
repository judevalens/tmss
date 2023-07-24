import os
import subprocess

def generate_proto_files():
    current_directory = os.getcwd()+"/proto/"
    print(current_directory)
    proto_files = [file for file in os.listdir(current_directory) if file.endswith(".proto")]

    if not proto_files:
        print("No proto files found in the current directory.")
        return

    for proto_file in proto_files:
        proto_path = os.path.join(current_directory, proto_file)
        command = f"protoc --proto_path={current_directory}  --go_out=gen {proto_path}"

        try:
            subprocess.run(command, shell=True, check=True)
            print(f"Generated {proto_file.replace('.proto', '_pb2.py')}")
        except subprocess.CalledProcessError as e:
            print(f"Failed to generate {proto_file}: {e}")

if __name__ == "__main__":
    generate_proto_files()