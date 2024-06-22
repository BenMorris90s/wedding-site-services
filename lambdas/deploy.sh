GUEST_INFO_TABLE_NAME=$(aws ssm get-parameter --name "/database/guest_info/table_name" --with-decryption --query "Parameter.Value" --output text)

sam build
sam deploy --parameter-overrides GuestInfoTableName="${GUEST_INFO_TABLE_NAME}"
