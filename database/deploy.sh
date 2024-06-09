GUEST_INFO_TABLE_NAME="GuestInfo"


sam build
sam deploy --parameter-overrides GuestInfoTableName="${GUEST_INFO_TABLE_NAME}"
# CF does not allow creation of a SecureString, so create it here instead.
aws ssm put-parameter --name "/database/guest_info/table_name" --value "${GUEST_INFO_TABLE_NAME}" --type "SecureString" --overwrite
