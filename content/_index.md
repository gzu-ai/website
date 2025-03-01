---
# Leave the homepage title empty to use the site title
title:
date: 2022-10-24
type: landing



sections:

  - block: slider
    content:
      slides:
      - title: 👋 Welcome to the group 🥳
        content: ''
        align: center
        background:
          image:
            filename: welcome1.png
            filters:
              brightness: 0.7
          position: right
          color: '#666'
      - title: Lecture & Learn 
        content: 'The research of the OKRR lab ranges from knowledge representation and reasoning to artificial neural networks!'
        align: left
        background:
          image:
            filename: welcome2.png
            filters:
              brightness: 0.7
          position: center
          color: '#555'
      - title: Visit & Lunch ☕️
        content: ' Welcome to join us! :smile:'
        align: left
        background:
          image:
            filename: welcome3.png
            filters:
              brightness: 0.7
          position: center
          color: '#555'
      - title: Seminar
        content: 'Team members have won awards such as the Ray Reiter Best Paper Award at KR-2006 and the Best Student Paper at ILP-2016.'
        align: left
        background:
          image:
            filename: welcome4.png
            filters:
              brightness: 0.7
          position: center
          color: '#555'
      - title: Omiga-krr Group
        content: 'Professor Wang Yisong has supervised more than 200 Ph.D. and M.S. students.🎓'
        align: right
        background:
          image:
            filename: welcome5.png
            filters:
              brightness: 0.5
          position: center
          color: '#333'
        link:
          icon: graduation-cap
          icon_pack: fas
          text: Join Us
          url: ../contact/
    design:
      # Slide height is automatic unless you force a specific height (e.g. '400px')
      slide_height: ''
      is_fullscreen: true
      # Automatically transition through slides?
      loop: false
      # Duration of transition between slides (in ms)
      interval: 2000

  
        #Located beside the picturesque Huaxi Park in Guiyang, China, the Omega-Knowledge Representation and Reasoning (OKRR) Laboratory is led by Professor Wang Yisong from Guizhou University. The laboratory focuses on cutting-edge research in computer software and theory, particularly in knowledge representation and reasoning, answer set programming, nonmonotonic reasoning, artificial intelligence (AI), and related fields such as machine learning. The lab has cultivated a dynamic and innovative team dedicated to advancing AI and knowledge engineering.

        #OKRR Laboratory’s research spans key areas of AI, including knowledge representation and reasoning, nonmonotonic reasoning, answer set programming, deep learning, and reinforcement learning.

        #Professor Wang Yisong and the team have authored the monograph Default Logic and Answer Set Programs (Science Press, 2023) and published numerous high-quality papers in internationally renowned journals and conferences, such as IEEE Transactions on Fuzzy Systems, Journal of Artificial Intelligence Research, ACM Transactions on Computational Logic, and Theory and Practice of Logic Programming.

        #The laboratory actively participates in top-tier international conferences, including the International Joint Conference on Artificial Intelligence (IJCAI), the Conference on Knowledge Representation and Reasoning (KR), the International Conference on Inductive Logic Programming (ILP), and the Conference on Logic Programming and Nonmonotonic Reasoning (LPNMR). Team members have garnered significant attention for their research achievements, winning honors such as the Ray Reiter Best Paper Award at KR-2006 and the Best Paper Award at ILP-2016.

        #In recent years, Professor Wang and his team have secured funding from the National Natural Science Foundation of China (General and Regional Programs), the Institute of Software at the Chinese Academy of Sciences, and the Guizhou Provincial Outstanding Scientific and Educational Talent Program. These grants underscore the team’s scientific excellence and potential.

        #OKRR Laboratory has nurtured over 200 doctoral and master’s graduates in the past decade. Students have received prestigious awards, including the National Scholarship, university-level academic scholarships, and prizes in the China Graduate Mathematical Modeling Competition and the National College Mathematical Modeling Contest. Many alumni have pursued doctoral studies at renowned institutions such as the University of Lyon (France), Vrije Universiteit Amsterdam (Netherlands), Tianjin University, Huazhong University of Science and Technology, and Wuhan University. Others contribute to local development in Guizhou, serving as key members at universities, Yunshang Guizhou, and the Guizhou Electric Power Design Institute.

        #The laboratory fosters a positive and collaborative team atmosphere, emphasizing innovation and cooperation. It maintains strong academic ties with leading institutions worldwide, including the Hong Kong University of Science and Technology, the University of Alberta (Canada), Griffith University, Western Sydney University (Australia), Vienna University of Technology (Austria), the University of Potsdam (Germany), Texas Tech University (USA), the Chinese Academy of Sciences, and the University of Science and Technology of China. These partnerships provide a broad platform for research and growth opportunities.

        #Moving forward, OKRR Laboratory remains committed to breakthroughs in knowledge representation, reasoning, AI, and logic programming, both theoretically and practically. We warmly welcome scholars and students passionate about these fields to join us in exploring technological frontiers and advancing the progress of artificial intelligence!


  
  - block: collection
    content:
      title: Latest News
      subtitle:
      text:
      count: 5
      filters:
        author: ''
        category: ''
        exclude_featured: false
        publication_type: ''
        tag: ''
      offset: 0
      order: desc
      page_type: post
    design:
      view: card
      columns: '1'
  

  - block: collection
    content:
      allLang: true
      title: Latest Publication
      text: ""
      count: 5
      filters:
        folders:
          - publication
        # publication_type: 'article'
    design:
      view: citation
      columns: '1'

  - block: markdown
    content:
      title:
      subtitle:
      text: |
        {{% cta cta_link="./people/" cta_text="Meet the team →" %}}
    design:
      columns: '1'



---
