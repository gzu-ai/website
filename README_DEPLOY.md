# 网站服务器部署说明

网站服务器使用 Nginx 提供 Hugo 生成的静态文件：

- 源码仓库：`/home/mr/website`
- 构建临时目录：`/home/mr/site-build`
- 线上目录：`/data/html/dev`
- Git 分支：`dev`
- Hugo：`/snap/bin/hugo`（Extended 版本）

> 不要把服务器密码、GitHub Token 或 SSH 私钥写入仓库。部署前应先确认需要发布的修改已经推送到 GitHub 的 `dev` 分支。

### 1. 登录服务器

```bash
ssh -p 22 mr@119.45.23.41
```

### 2. 更新服务器源码

```bash
set -euo pipefail

cd /home/mr/website

# 仓库应保持干净；如果这里出现文件，不要直接覆盖，先确认改动来源。
git status --short
test -z "$(git status --porcelain)"

git fetch origin dev
git checkout dev
git pull --ff-only origin dev

# 记录本次部署的提交，便于排查和回滚。
git log -1 --oneline
```

### 3. 在临时目录构建

主题模块 `blox-core` 和 `blox-seo` 的已发布版本存在模块路径拼写不一致问题，因此不要直接修改源码仓库的 `go.mod`。复制一个独立的构建副本，并在副本中添加本地模块映射：

```bash
set -euo pipefail

BUILD_SOURCE=/home/mr/website-build-src
BUILD_OUTPUT=/home/mr/site-build
MODULE_ROOT=/home/mr/hugo-modules

rm -rf /home/mr/website-build-src /home/mr/site-build
cp -a /home/mr/website "$BUILD_SOURCE"

cd "$BUILD_SOURCE"

go mod edit \
  -replace=github.com/gzu-ai/hugo-blox-builder/modules/blox-core="$MODULE_ROOT/core"

go mod edit \
  -replace=github.com/gzu-ai/hugo-blox-builder/modules/blox-seo="$MODULE_ROOT/seo"

mkdir -p "$BUILD_OUTPUT"

HUGO_MODULE_PROXY=https://goproxy.cn,direct \
GOPROXY=https://goproxy.cn,direct \
/snap/bin/hugo --minify --destination "$BUILD_OUTPUT"
```

构建成功时会输出中英文页面数量。出现 `Error` 时不要继续发布；Hugo 的 `deprecated` 警告暂不影响构建。

### 4. 检查构建结果

```bash
test -f /home/mr/site-build/zh/index.html
test -f /home/mr/site-build/en/index.html

du -sh /home/mr/site-build
```

如果修改了具体页面或附件，还应在临时目录检查对应文件。例如：

```bash
# 检查新闻页面
test -f /home/mr/site-build/zh/post/index.html

# 检查王以松老师的 CV
test -f /home/mr/site-build/zh/resume/cv-yisong.pdf
test -f /home/mr/site-build/en/resume/cv-yisong.pdf
```

### 5. 备份并发布

下面的操作先把新站点同步到独立目录，检查 Nginx 配置后再切换线上目录。旧站点会保留为带时间戳的备份。

```bash
set -euo pipefail

STAMP=$(date +%Y%m%d-%H%M%S)
STAGE="/data/html/dev.new.$STAMP"
BACKUP="/data/html/dev.backup.$STAMP"

sudo mkdir "$STAGE"
sudo rsync -a --delete --chown=root:root /home/mr/site-build/ "$STAGE/"
sudo nginx -t

sudo mv /data/html/dev "$BACKUP"

if ! sudo mv "$STAGE" /data/html/dev; then
  # 新目录切换失败时立即恢复旧站点。
  sudo mv "$BACKUP" /data/html/dev
  exit 1
fi

echo "部署完成"
echo "旧站备份：$BACKUP"
```

Nginx 直接读取静态文件，正常情况下不需要重启。

### 6. 验证网站

先从服务器本机验证：

```bash
curl -I http://127.0.0.1/
curl -I http://127.0.0.1/zh/
curl -I http://127.0.0.1/en/
```

再通过浏览器或另一台电脑访问：

```text
http://119.45.23.41/
```

### 回滚到上一次部署

如果新版本出现问题，找到最近的备份目录后执行：

```bash
ls -dt /data/html/dev.backup.* | head
```

确认要恢复的目录后，将下面的 `<备份目录>` 替换为完整路径：

```bash
set -euo pipefail

FAILED="/data/html/dev.failed.$(date +%Y%m%d-%H%M%S)"

sudo mv /data/html/dev "$FAILED"
sudo mv <备份目录> /data/html/dev

curl -I http://127.0.0.1/
```
