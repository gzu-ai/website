# 成员资料管理后台配置

本项目在 `/admin/` 提供 Decap CMS 成员资料编辑器。编辑器使用 GitHub 登录、Open Authoring 和 Editorial Workflow：没有仓库写权限的同学会在自己的 Fork 中保存修改，并通过 Pull Request 请求管理员审核；合并到 `dev` 后由 GitHub Actions 自动部署。

## 上线前置条件

1. 为网站配置正式域名及 HTTPS。管理后台会拒绝在公网 HTTP 页面加载。
2. 在 GitHub 创建 OAuth App：
   - Homepage URL：`https://www.omega-krr-gz.cn/admin/`
   - Authorization callback URL：`https://www.omega-krr-gz.cn/callback`
   - 不要启用 callback URL 通配符。
3. 将 Client ID 和 Client Secret 仅保存到服务器 `/etc/omega-cms-oauth.env`，不要提交到 Git。

环境文件示例：

```dotenv
GITHUB_CLIENT_ID=<GitHub OAuth App Client ID>
GITHUB_CLIENT_SECRET=<GitHub OAuth App Client Secret>
OAUTH_PUBLIC_URL=https://www.omega-krr-gz.cn
CMS_ALLOWED_ORIGIN=https://www.omega-krr-gz.cn
OAUTH_LISTEN=127.0.0.1:3000
GITHUB_SCOPE=public_repo user:email
```

## 安装 OAuth 代理

在服务器源码目录执行：

```bash
cd /home/mr/website/tools/cms-oauth
go test ./...
go build -o /tmp/omega-cms-oauth .

sudo install -o root -g root -m 0755 /tmp/omega-cms-oauth /usr/local/bin/omega-cms-oauth
sudo install -o root -g root -m 0644 /home/mr/website/deploy/server/cms-oauth.service /etc/systemd/system/cms-oauth.service
sudo chown root:root /etc/omega-cms-oauth.env
sudo chmod 0600 /etc/omega-cms-oauth.env
sudo systemctl daemon-reload
sudo systemctl enable --now cms-oauth.service
curl http://127.0.0.1:3000/healthz
```

将 `deploy/server/nginx-cms.conf.example` 中的两个 `location` 加入网站 HTTPS server block，然后执行：

```bash
sudo nginx -t
sudo systemctl reload nginx
```

HTTPS 证书续期由 `certbot-renew.timer` 每日自动检查。

## 自动部署配置

服务器上安装受限的发布脚本：

```bash
sudo install -o root -g root -m 0755 \
  /home/mr/website/deploy/server/deploy-site \
  /usr/local/sbin/deploy-lab-site

echo 'mr ALL=(root) NOPASSWD: /usr/local/sbin/deploy-lab-site *' \
  | sudo tee /etc/sudoers.d/omega-website-deploy
sudo chmod 0440 /etc/sudoers.d/omega-website-deploy
sudo visudo -cf /etc/sudoers.d/omega-website-deploy
```

生成专用部署密钥；私钥只保存为 GitHub Actions Secret，公钥放入服务器 `mr` 用户的 `authorized_keys`：

```bash
ssh-keygen -t ed25519 -C 'gzu-ai website deploy' -f ./gzu-ai-website-deploy
```

在 GitHub 仓库的 **Settings → Secrets and variables → Actions** 添加：

- `DEPLOY_HOST`：服务器域名或 IP
- `DEPLOY_PORT`：`22`
- `DEPLOY_USER`：`mr`
- `DEPLOY_SSH_KEY`：部署专用私钥全文
- `DEPLOY_KNOWN_HOSTS`：`ssh-keyscan -H <服务器域名或 IP>` 的输出

## 分支保护

在 GitHub 的 `dev` 分支规则中启用：

- Require a pull request before merging
- Require approvals（至少 1 人）
- Require status checks，选择 `Build and deploy website / build`
- Block force pushes
- Restrict deletions

当前 Code Owner 是已经确认的 `@2hlovely`。如需多人审核，可在 `.github/CODEOWNERS` 中加入其他实验室管理员的 GitHub 用户名或团队。

## 使用方式

1. 同学访问 `https://www.omega-krr-gz.cn/admin/` 并通过 GitHub 登录。
2. 新建或修改自己的中英文资料，保存草稿后选择“提交审核”。
3. GitHub 自动创建 Pull Request；管理员检查内容与构建结果后合并。
4. 合并到 `dev` 后 GitHub Actions 自动构建、备份并发布。

Open Authoring 允许没有仓库写权限的 GitHub 用户提交修改，但它不能严格限制用户只能改自己的文件。管理员必须审核每个 PR；若需要强制逐人权限，需要额外开发账号与作者目录映射服务。
