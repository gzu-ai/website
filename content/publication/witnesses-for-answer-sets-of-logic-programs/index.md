---
title: "Witnesses for Answer Sets of Logic Programs"
authors:
- Yisong Wang
- Thomas Eiter
- Yuanlin Zhang
- Fangzhen Lin
# author_notes:
# - "Equal contribution"
# - "Equal contribution"
date: "2023-04-01T00:00:00Z"
doi: "10.1145/3568955"

# Schedule page publish date (NOT publication's date).
publishDate: "2023-01-27T00:00:00Z"

# Publication type.
# Accepts a single type but formatted as a YAML list (for Hugo requirements).
# Enter a publication type from the CSL standard.
publication_types: ["article-journal"]

# Publication name and optional abbreviated publication name.
publication: "ACM Trans. Comput. Logic"
publication_short: "ACM Trans. Comput. Logic"

abstract: In this article, we consider Answer Set Programming (ASP). It is a declarative problem solving paradigm that can be used to encode a problem as a logic program whose answer sets correspond to the solutions of the problem. It has been widely applied in various domains in AI and beyond. Given that answer sets are supposed to yield solutions to the original problem, the question of “why a set of atoms is an answer set” becomes important for both semantics understanding and program debugging. It has been well investigated for normal logic programs. However, for the class of disjunctive logic programs, which is a substantial extension of that of normal logic programs, this question has not been addressed much. In this article, we propose a notion of reduct for disjunctive logic programs and show how it can provide answers to the aforementioned question. First, we show that for each answer set, its reduct provides a resolution proof for each atom in it. We then further consider minimal sets of rules that will be sufficient to provide resolution proofs for sets of atoms. Such sets of rules will be called witnesses and are the focus of this article. We study complexity issues of computing various witnesses and provide algorithms for computing them. In particular, we show that the problem is tractable for normal and headcycle-free disjunctive logic programs, but intractable for general disjunctive logic programs. We also conducted some experiments and found that for many well-known ASP and SAT benchmarks, computing a minimal witness for an atom of an answer set is often feasible.

# Summary. An optional shortened abstract.
# summary: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis posuere tellus ac convallis placerat. Proin tincidunt magna sed ex sollicitudin condimentum.

tags:
- ASP
- Witness
featured: false

# links:
# - name: ""
#   url: ""
# url_pdf: http://arxiv.org/pdf/1512.04133v1
url_code: 'https://github.com/yswang168/witness'
# url_dataset: ''
# url_poster: ''
# url_project: ''
# url_slides: ''
# url_source: ''
# url_video: ''

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder. 
image:
  caption: ''
  focal_point: ""
  preview_only: false

# Associated Projects (optional).
#   Associate this publication with one or more of your projects.
#   Simply enter your project's folder or file name without extension.
#   E.g. `internal-project` references `content/project/internal-project/index.md`.
#   Otherwise, set `projects: []`.
projects: []

# Slides (optional).
#   Associate this publication with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides: "example"` references `content/slides/example/index.md`.
#   Otherwise, set `slides: ""`.
# slides: example
---