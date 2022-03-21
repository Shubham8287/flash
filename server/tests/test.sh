DATA_PATH='../src/data/words_dictionary.json'
TEST_DATA_PATH='./test_data.json'

echo "[" > $TEST_DATA_PATH
totalLines=$(wc -l < $DATA_PATH)

for i in `eval echo {2.."$1"}`
do
    pickedLine=$(sed -n "$((($RANDOM*$RANDOM)%$totalLines))"p $DATA_PATH)
    echo $pickedLine  >> $TEST_DATA_PATH 
done

pickedLine=$(sed -n "$((($RANDOM*$RANDOM)%$totalLines))"p $DATA_PATH)
echo $pickedLine | sed -e 's/\",/"/g'  >> $TEST_DATA_PATH

echo "]" >> $TEST_DATA_PATH