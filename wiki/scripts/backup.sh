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
  docker exec --user 0 wiki cp -r ${CONF_DIR} ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/pages ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/attic ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/meta ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/meta_attic ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/media ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/media_meta ${BACKUPS_DIR}
  docker exec --user 0 wiki cp -r ${DATA_DIR}/media_attic ${BACKUPS_DIR}
  docker cp ${container_id}:${BACKUPS_DIR} ..${BACKUPS_DIR}
  echo "backup finished"
elif [[ ${arg} == "restore" ]]; then
  docker cp ..${BACKUPS_DIR}/. ${container_id}:${RESTORE_DIR}
  docker exec --user 1001 wiki cp -r ${RESTORE_DIR}/conf /bitnami/dokuwiki
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/pages ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/attic ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/meta ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/meta_attic ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/media ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/media_meta ${DATA_DIR}
  docker exec --user 1001 wiki cp -r $RESTORE_DIR/media_attic ${DATA_DIR}
  echo "restore finished."
else 
  echo "The script needs an argument ('backup' or 'restore')"
fi