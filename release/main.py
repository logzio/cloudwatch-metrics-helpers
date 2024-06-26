import os

import boto3

REGIONS = [
    'us-east-1', 'us-east-2', 'us-west-1', 'us-west-2',
    'ap-south-1', 'ap-northeast-3', 'ap-northeast-2', 'ap-southeast-1', 'ap-southeast-2', 'ap-northeast-1',
    'eu-central-1', 'eu-west-1', 'eu-west-2', 'eu-west-3', 'eu-north-1',
    'sa-east-1',
    'ca-central-1',
    'eu-central-2', 'eu-south-1', 'eu-south-2', 'ap-south-2', 'ap-southeast-3', 'ap-southeast-4', 'ap-east-1', 'ca-west-1', 'af-south-1', 'me-central-1', 
    'il-central-1','me-south-1'
]

BUCKET_NAME_PREFIX = 'logzio-aws-integrations-'
ENV_ACCESS_KEY = 'AWS_ACCESS_KEY'
ENV_SECRET_KEY = 'AWS_SECRET_KEY'
ENV_FOLDER_NAME = 'FOLDER_NAME'
ENV_VERSION_NUMBER = 'VERSION_NUMBER'
ENV_PATH_TO_FILE = 'PATH_TO_FILE'
CF_TEMPLATE = 'sam-template.yaml'
CF_TEMPLATE_S3 = 'sam-s3-daily-metrics.yaml'
REGION_PLACEHOLDER = '<<REGION>>'
VERSION_PLACEHOLDER = '<<VERSION>>'


def upload_public_to_s3(access_key, secret_key, folder_name, version_number, path_to_file):
    file_name = path_to_file.split('/')[-1]
    print(f'File name: {file_name}')
    success = 0
    for region in REGIONS:
        try:
            print(f'Region: {region}')
            # Create a new S3 client for each region
            s3 = get_s3_client(access_key, secret_key, region)
            object_name = f'{folder_name}/{version_number}/{file_name}'
            bucket_name = f'{BUCKET_NAME_PREFIX}{region}'
            s3.upload_file(path_to_file, bucket_name, object_name, ExtraArgs={'ACL': 'public-read'})
            success += 1
        except Exception as e:
            print(f'Error occurred for region {region}: {e}')
            print('Skipping this region')

    print(f'Uploaded to {success} regions')


def cf_template_workflow(access_key, secret_key, folder_name, version_number, path_to_file):
    file_name = path_to_file.split('/')[-1]
    print(f'File name: {file_name}')
    success = 0
    base_arr = []
    with open(path_to_file, 'r') as base_file:
        base_arr = base_file.readlines()
    if len(base_arr) == 0:
        raise ValueError('Could not get base Cloudformation template')
    for region in REGIONS:
        try:
            print(f'Region: {region}')
            print(f'Version: {version_number}')
            # Adjust the template content for each region
            tmp_arr = [line.replace(REGION_PLACEHOLDER, region).replace(VERSION_PLACEHOLDER, version_number) for line in base_arr]
            new_path = f'./{file_name}' 
            with open(new_path, 'w') as new_file:
                new_file.writelines(tmp_arr)
            object_name = f'{folder_name}/{version_number}/{file_name}'
            bucket_name = f'{BUCKET_NAME_PREFIX}{region}'
            # Create a new S3 client for the current region
            s3 = get_s3_client(access_key, secret_key, region)
            s3.upload_file(new_path, bucket_name, object_name, ExtraArgs={'ACL': 'public-read'})
            success += 1
            os.remove(new_path)  # Clean up the locally created file
        except Exception as e:
            print(f'Error occurred for region {region}: {e}')
            print('Skipping this region')

    print(f'Uploaded to {success} regions')

def get_s3_client(access_key, secret_key, region_name):
    session = boto3.Session(
        aws_access_key_id=access_key,
        aws_secret_access_key=secret_key,
        region_name=region_name
    )

    return session.client('s3')


def upload():
    access_key = os.getenv(ENV_ACCESS_KEY)
    secret_key = os.getenv(ENV_SECRET_KEY)
    if access_key is None or access_key == '' or secret_key is None or secret_key == '':
        raise ValueError('AWS credentials missing! Exiting')
    folder_name = os.getenv(ENV_FOLDER_NAME)
    if folder_name is None or folder_name == '':
        raise ValueError('Missing folder name! Exiting')
    version_number = os.getenv(ENV_VERSION_NUMBER)
    if version_number is None or version_number == '':
        raise ValueError('Missing version number! Exiting')
    path_to_file = os.getenv(ENV_PATH_TO_FILE)
    if path_to_file is None or path_to_file == '':
        raise ValueError('Missing path to file! Exiting')
    file_exists = os.path.isfile(path_to_file)
    if not file_exists:
        raise FileNotFoundError(f'Provided path to file ({path_to_file}) does not exists! Exiting')
    try:
        is_cf_template = (path_to_file.split('/')[-1] == CF_TEMPLATE or path_to_file.split('/')[-1] == CF_TEMPLATE_S3)
        print(f'Is Cloudformation template: {is_cf_template}')
        if is_cf_template:
            cf_template_workflow(access_key, secret_key, folder_name, version_number, path_to_file)
        else:
            upload_public_to_s3(access_key, secret_key, folder_name, version_number, path_to_file)
    except Exception as e:
        print(f'Some error occurred while trying to upload file: {e}')


if __name__ == '__main__':
    upload()
