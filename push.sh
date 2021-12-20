echo
echo "-------------------- Tools 仓库"
## 执行上传
git pull
git add -A .
git commit -m 'auto'
ver=`tail -n 1 CHANGELOG.md | awk '{print $NF}'`
# git tag -f "v1.0.1"
git tag -f "v${ver}"
git push --tags --force
git push

