import os
import shutil

def move_mp4_to_individual_folders(root_dir):
    for entry in os.scandir(root_dir):
        print(f"Found dir {entry.name}")
        if not entry.is_dir():
            continue
        for filename in os.listdir(entry):
            if filename.endswith('.mp4'):
                print(f"Found mp4 {filename}")
                video_name = os.path.splitext(filename)[0]
                video_folder_path = os.path.join(entry, video_name)

                if not os.path.exists(video_folder_path):
                    os.makedirs(video_folder_path)

                video_file_path = os.path.join(entry.path, filename)
                new_video_file_path = os.path.join(video_folder_path, filename)
                print(f"Moving {video_file_path} to {new_video_file_path}")

                shutil.move(video_file_path, new_video_file_path)

if __name__ == "__main__":
    print("Starting to move mp4 files to individual folders")
    root_directory = "public/videos"  # Replace with the path to your 'videos' directory
    move_mp4_to_individual_folders(root_directory)
