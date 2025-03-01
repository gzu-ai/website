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

  
        
        #Omega-Knowledge Representation and Reasoning(OKRR)实验室坐落在美丽的中国贵阳市花溪公园旁，由贵州大学王以松教授(学术带头人、博士生导师)组建领导，专注于计算机软件与理论中的知识表示与推理、回答集程序设计、非单调推理等人工智能以及相关机器学习领域的前沿研究。实验室凝聚了一支充满活力和创新精神的团队，致力于推动人工智能和知识工程相关领域的发展。OKRR实验室的研究领域涵盖了知识表示与推理、非单调推理、回答集编程、深度学习、强化学习等人工智能的重要方面。
        
        #王以松教授及实验室团队出版了专著《Default Logic and Answer Set Programs》(科学出版社, 2023)，并在多种国际知名期刊和会议上发表了多篇高水平论文，包括但不限于《IEEE Transactions on Fuzzy Systems》、《Journal of Artificial Intelligence Research》、《ACM Transactions on Computational Logic》、《Theory and Practice of Logic Programming》等。
        
        #在会议论文方面，我们积极参与了国际顶级会议的交流与展示，如国际人工智能联合会议 (IJCAI)、知识表示会议 (KR)、归纳逻辑程序设计会议(ILP)逻辑编程和非单调推理会议 (LPNMR) 等。团队成员的研究成果屡次在这些顶级会议上获得广泛关注，并在KR-2006大会上获得Ray Reiter最佳论文，在ILP-2016大会上获得最佳论文等荣誉。

        #在科研项目方面，王教授及其团队近年来获得了包括国家自然科学基金（面上项目和地区项目）、中科院软件所资助以及贵州省优秀科技教育人才项目在内的多项资助。这些支持不仅为我们的研究提供了坚实的资金基础，也进一步证明了我们团队在科研领域的卓越贡献和潜力。
        
        #在人才培养方面，OKRR实验室培养的学生多次获得国家奖学金、校级学业奖学金，以及研究生数学建模竞赛和全国高校数学建模竞赛的奖项。我们致力于为学生提供扎实的学术基础与丰富的实践机会。近十年来，实验室累计培养了博士、硕士研究生200余名。其中，多名硕士被法国里昂大学、荷兰自由大学、天津大学、华中科技大学、武汉大学等知名高校录取，继续深造攻读博士学位。此外，多名毕业生选择扎根贵州，活跃在贵州省各大高校、云上贵州、贵州电力设计院等单位，部分成为单位的核心骨干，为地方的科技进步和社会发展贡献了力量。
        
        #我们拥有积极良好的团队氛围，注重培养学生的合作精神与创新能力。实验室持续开展国内外的合作与学术交流，与香港科技大学、加拿大 Alberta 大学、澳大利亚 Griffith 大学及西悉尼大学、奥地利维也纳科技大学、德国 Potsdam 大学、美国德州理工大学、中科院、中国科学技术大学等多所国内外知名学府保持紧密联系。这种广泛的合作交流不仅为实验室的科研提供了更广阔的平台，也为学生们创造了更多的学习和成长机会。
        
        #我们未来将继续致力于知识表示与推理、人工智能以及逻辑编程领域的研究，一直追寻在理论和应用上都能持续取得突破。我们欢迎对相关研究方向感兴趣的学者和学生加入我们，共同探索科技前沿，共同为人工智能的进步做出贡献！
  
  - block: collection
    content:
      title: 新闻
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
      title: 最新发表
      allLang: true
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
        {{% cta cta_link="./people/" cta_text="团队成员 →" %}}
    design:
      columns: '1'
---
