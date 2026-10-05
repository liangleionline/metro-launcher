# Metro Launcher — Windows 8「开始」屏幕风格导航首页

在浏览器里复现 Windows 8 的「开始」屏幕，把它变成你的**浏览器首页 / NAS 导航页 / 电视遥控器导航页**。

视觉实现基于 [GeeLaw](https://gist.github.com/GeeLaw/bc61a3f724d45b9e9df9faf8044312d8) 的开源作品（MIT），在其之上做了导航化改造：

- ✅ 磁贴由配置数组驱动，每个磁贴可指向任意链接（NAS 服务、电视应用、常用网站）
- ✅ 点击磁贴：放大过渡动画后在新标签页打开目标链接
- ✅ 键盘 / 电视遥控器方向键导航（←↑↓→ 移动焦点，Enter 打开）
- ✅ 响应式缩放，适配电脑浏览器与电视浏览器（含 4K）
- ✅ 顶部实时时钟 + 用户名，还原 Win8 桌面氛围
- ✅ 保留原版细节：深紫背景、视差滚动、磁贴 shine 渐变、按点击位置变化的按压放大动画
- ✅ 打包为飞牛 fnOS 的 `.fpk` 应用，安装后在 `:5080` 端口访问

## 快速使用

1. 打开 `index.html`（或部署到任意 Web 服务器 / NAS）。
2. 编辑文件顶部 `CONFIG` 配置区：
   - `userName` / `userAvatar`：右上角显示的用户名与头像；
   - `tiles`：磁贴列表，每项 `{ title, url, color, icon, wide }`；
   - 把 `url` 里的 `<NAS-IP>` 替换成你飞牛 NAS 的真实地址即可。

```js
tiles: [
  { title: '控制台', url: 'http://<NAS-IP>:5666', color: '#0067b8', icon: '🖥️', wide: false },
  { title: '哔哩哔哩', url: 'https://www.bilibili.com', color: '#fb7299', icon: '📺', wide: true }
]
```

### 设为浏览器首页
- Chrome / Edge：设置 → 启动时 → 特定网页，填你的部署地址；
- 电视浏览器（如 TV 端浏览器）：打开该地址并设为首页，用遥控器方向键在磁贴间移动、OK 键打开。

## 部署到飞牛 fnOS（fpk 应用）

```bash
cd metro-launcher
# 若未安装打包工具，先下载 fnpack：
#   https://developer.fnnas.com/docs/cli/fnpack
fnpack build        # 生成 metro-launcher.fpk（位于 metro-launcher/ 目录）
```

1. 打开飞牛「应用中心」→ 手动安装，上传生成的 `.fpk`；
2. 安装完成后，从桌面图标或浏览器访问 `http://<NAS-IP>:5080`。

> 应用基于 `nginx:alpine` 容器，监听宿主 5080 端口，把 `app/docker/web/index.html` 挂载进容器。

## 目录结构

```
├── index.html               # 单文件导航首页（浏览器直接打开）
├── metro-launcher/          # 飞牛 fnOS fpk 应用工程
│   ├── app/docker/          # Docker 编排 + web 静态页
│   ├── config/              # 权限与资源声明
│   ├── cmd/                 # 生命周期脚本
│   └── manifest             # 应用元信息
└── LICENSE
```

## 致谢与许可

磁贴视觉基于 [GeeLaw 的 tile-animation.html](https://gist.github.com/GeeLaw/bc61a3f724d45b9e9df9faf8044312d8)（Copyright © 2025 Ji Luo），本项目在其上修改并扩展。均以 [MIT](./LICENSE) 协议分发。
