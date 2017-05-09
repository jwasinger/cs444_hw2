folder=./tmp
mkdir $folder
trap "rm -rf $folder && exit 0" SIGINT SIGTERM #remove temp files when exiting

while :
do
  # create a file named f-{{current epoch time}} with the contents "hello world"
  x=$(date +%s)
  touch $folder/f-$x && echo "hello world" >> $folder/f-$x
  sleep 1s
done
