(function () {
  "use strict";

  const translatedString = (label, name, required = true) => ({
    label,
    name,
    widget: "string",
    required,
    i18n: true,
  });

  const config = {
    load_config_file: false,
    locale: "zh_Hans",
    site_url: window.location.origin,
    display_url: window.location.origin,
    logo_url: "/media/logo.png",
    backend: {
      name: "github",
      repo: "gzu-ai/website",
      branch: "dev",
      base_url: window.location.origin,
      auth_endpoint: "/auth",
      site_domain: window.location.hostname,
      open_authoring: true,
      squash_merges: true,
      use_graphql: true,
    },
    publish_mode: "editorial_workflow",
    media_folder: "static/media/uploads",
    public_folder: "/media/uploads",
    i18n: {
      structure: "multiple_files",
      locales: ["zh", "en"],
      default_locale: "zh",
    },
    collections: [
      {
        name: "authors",
        label: "成员资料",
        label_singular: "成员资料",
        description: "编辑中英文个人资料。保存后会创建草稿；提交审核后由管理员合并 Pull Request。",
        folder: "content/authors",
        create: true,
        extension: "md",
        format: "frontmatter",
        i18n: true,
        nested: {
          depth: 2,
          summary: "{{title}}",
        },
        meta: {
          path: {
            label: "个人页面目录名",
            widget: "string",
            hint: "首次创建时填写小写英文名，例如 yuxinzhong；创建后不要修改。",
            index_file: "_index",
          },
        },
        media_folder: "",
        public_folder: "",
        fields: [
          translatedString("姓名", "title"),
          {
            label: "作者标识",
            name: "authors",
            widget: "list",
            required: true,
            i18n: true,
            field: { label: "标识", name: "author", widget: "string" },
            hint: "通常填写个人页面目录名；如已有论文作者标识，请保留原值。",
          },
          { label: "主站负责人", name: "superuser", widget: "boolean", default: false, i18n: "duplicate" },
          translatedString("身份/职位", "role"),
          {
            label: "单位",
            name: "organizations",
            widget: "list",
            required: false,
            i18n: true,
            fields: [
              { label: "单位名称", name: "name", widget: "string" },
              { label: "单位网址", name: "url", widget: "string", required: false },
            ],
          },
          { label: "简短简介", name: "bio", widget: "text", required: false, i18n: true },
          {
            label: "研究兴趣",
            name: "interests",
            widget: "list",
            required: false,
            i18n: true,
            field: { label: "研究方向", name: "interest", widget: "string" },
          },
          {
            label: "教育经历",
            name: "education",
            widget: "object",
            required: false,
            i18n: true,
            fields: [
              {
                label: "经历",
                name: "courses",
                widget: "list",
                required: false,
                fields: [
                  { label: "学位/专业", name: "course", widget: "string" },
                  { label: "学校", name: "institution", widget: "string" },
                  { label: "时间", name: "year", widget: "string", required: false },
                ],
              },
            ],
          },
          {
            label: "头像",
            name: "avatar",
            widget: "image",
            required: false,
            i18n: "duplicate",
            media_library: { config: { max_file_size: 5242880 } },
            hint: "建议上传正方形 JPG/PNG/WebP，最大 5 MB。",
          },
          {
            label: "联系方式与个人主页",
            name: "social",
            widget: "list",
            required: false,
            i18n: true,
            fields: [
              {
                label: "图标",
                name: "icon",
                widget: "select",
                options: ["envelope", "github", "linkedin", "google-scholar", "cv", "globe"],
              },
              {
                label: "图标库",
                name: "icon_pack",
                widget: "select",
                options: ["fas", "fab", "ai"],
                default: "fas",
              },
              { label: "链接", name: "link", widget: "string" },
              { label: "说明", name: "label", widget: "string", required: false },
            ],
          },
          { label: "Gravatar 邮箱", name: "email", widget: "string", required: false, i18n: "duplicate" },
          { label: "突出显示姓名", name: "highlight_name", widget: "boolean", default: false, i18n: "duplicate" },
          {
            label: "成员分组",
            name: "user_groups",
            widget: "list",
            required: true,
            i18n: true,
            field: {
              label: "分组",
              name: "group",
              widget: "select",
              options: [
                "导师", "首席研究员", "研究员", "行政", "访问学者", "博士生", "硕士生", "博士毕业生", "硕士毕业生",
                "Supervisors", "Principal Investigators", "Researchers", "Administration", "Visitors", "PhD Candidate", "Grad Students", "PhD Alumni", "Master's Alumni"
              ],
            },
            hint: "中文页选择中文分组，英文页选择英文分组。",
          },
          { label: "详细介绍", name: "body", widget: "markdown", required: false, i18n: true },
        ],
      },
    ],
  };

  window.CMS.init({ config });
})();
