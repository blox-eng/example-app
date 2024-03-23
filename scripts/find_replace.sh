grep -rl 'service/handler' . | xargs sed -i '' 's/service\/handler/service\/internal\/handlers/g'

grep -rl '/app/' . | xargs sed -i '' 's/\/backend\//\/app\//g'

grep -rl '/model' . | xargs sed -i '' 's/\/model/\/model/g'

