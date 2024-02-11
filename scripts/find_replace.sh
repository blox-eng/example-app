grep -rl 'service/handler' . | xargs sed -i '' 's/service\/handler/service\/internal\/handlers/g'

grep -rl '/backend/' . | xargs sed -i '' 's/\/service\//\/backend\//g'

grep -rl '/model' . | xargs sed -i '' 's/\/model/\/model/g'

