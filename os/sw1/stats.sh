for i in $(find /); do
if [ -f $i ]; then
echo $(stat --format=%s "$i")
fi
done > stats.txt
