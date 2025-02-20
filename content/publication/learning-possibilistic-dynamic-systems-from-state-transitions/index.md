---
title: "Learning possibilistic dynamic systems from state transitions"
authors:
- hongbohu
- yisongwang
- Katsumi Inoue
# author_notes:
# - "Equal contribution"
# - "Equal contribution"
date: "2025-01-01T00:00:00Z"
doi: "/10.1016/j.fss.2024.109259"

# Schedule page publish date (NOT publication's date).
publishDate: "2025-01-01T00:00:00Z"

# Publication type.
# Accepts a single type but formatted as a YAML list (for Hugo requirements).
# Enter a publication type from the CSL standard.
publication_types: ["article-journal"]

# Publication name and optional abbreviated publication name.
publication: "Fuzzy Sets and Systems"
publication_short: "Fuzzy Sets and Systems"

abstract: Learning from 1-step transitions (LF1T) has become a paradigm to construct a logical hypothesis of a dynamic system, such as a Boolean network, from its synchronized state transitions and background knowledge. While uncertain and incomplete information plays an important role in dynamic systems, LF1T and its successors cannot handle uncertainty modeled by possibility theory. This motivates our combination of inductive logic programming (ILP) and possibilistic normal logic program (poss-NLP) that applies to reasoning about uncertain dynamic systems. In this paper, we propose a learning task to learn a poss-NLP from given interpretation transitions and background knowledge. The sufficient and necessary condition for the existence of its solution is determined. We introduce an algorithm called iltp to learn a specific solution, which typically encompasses mass redundant rules. Additionally, we propose another algorithm called sp-iltp to identify global minimal solutions. Alongside theoretical correctness proofs, a synthetic experiment demonstrates the learning performance on six gene regulatory networks with possibilistic uncertainty. This work thus offers a rational framework for learning the dynamics of systems under uncertainty via poss-NLPs.
# Summary. An optional shortened abstract.
# summary: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis posuere tellus ac convallis placerat. Proin tincidunt magna sed ex sollicitudin condimentum.

tags:
- Uncertainty
- Inductive learning
- Possibilistic logic programs
- Boolean networks
- Gene regulatory network
featured: false

# links:
# - name: ""
#   url: ""
# url_pdf: http://arxiv.org/pdf/1512.04133v1
url_code: 'https://github.com/gzu-ai/ILP'
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