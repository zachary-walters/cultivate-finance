arg=$1

BACKUPS_DIR=/cfbackups
RESTORE_DIR=/cfrestore

DATA_DIR=/bitnami/dokuwiki/data
CONF_DIR=/bitnami/dokuwiki/conf

container_id=$(docker ps -aqf "name=wiki")



if [[ ${arg} == "backup" ]]; then 
  echo "backing up wiki..."
  rm -rf ..${BACKUPS_DIR}
  docker exec --user 0 wiki rm -rf ${BACKUPS_DIR}
  docker exec --user 0 wiki mkdir -p ${BACKUPS_DIR}
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/conf.zip ${CONF_DIR}
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/pages.zip ${DATA_DIR}/pages
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/attic.zip ${DATA_DIR}/attic
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/meta.zip ${DATA_DIR}/meta
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/meta_attic.zip ${DATA_DIR}/meta_attic
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/media.zip ${DATA_DIR}/media
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/media_meta.zip ${DATA_DIR}/media_meta
  docker exec --user 0 wiki zip -r ${BACKUPS_DIR}/media_attic.zip ${DATA_DIR}/media_attic/
  docker cp ${container_id}:${BACKUPS_DIR} ..${BACKUPS_DIR}
  echo "backup finished"
elif [[ ${arg} == "restore" ]]; then
  docker cp ..${BACKUPS_DIR}/. ${container_id}:${RESTORE_DIR}
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/conf.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/pages.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/attic.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/meta.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/meta_attic.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/media.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/media_meta.zip -d / 
  docker exec --user 1001 wiki unzip -o ${RESTORE_DIR}/media_attic.zip -d / 
  echo "restore finished."
else 
  echo "The script needs an argument ('backup' or 'restore')"
fi