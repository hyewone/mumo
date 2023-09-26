# mumo


## 목차
- [프로젝트 소개](#project-introduction)
- [깃 레포지토리](#git-repo)
- [주요 기능](#main-function)
- [개발 환경](#development-environment)
- [CI/CD 파이프라인](#cicd-pipeline)

</br>

<a name="project-introduction"></a>
## 프로젝트 소개
메가박스, CGV, 롯데시네마 시네마 사이트에서 크롤링한 무대인사 정보를 지도, 달력 등 다양한 형태로 모아보는 토이 프로젝트입니다.

</br>

<a name="git-repo"></a>
## 깃 레포지토리
| 레포지토리 | 역할 | 
|----------|----------|
| [mumo](https://github.com/hyewone/mumo) | 크롤링, RestAPI, 인증, 인가, JWT 토큰 발행 |
| [mumo-react](https://github.com/hyewone/mumo-react) | 화면 |

</br>

<a name="main-function"></a>
## 주요 기능
- 시네마의 무대인사 정보 사이트 바로가기
  
<img src="images/화면 기록 2023-09-25 오후 3.15.38 (1).gif"></img>
</br>
</br>
- 무대인사 정보 지도로 보기
  
  
<img src="images/화면 기록 2023-09-25 오후 3.16.57.gif"></img>
</br>
</br>
- 무대인사 정보 달력으로 보기
  
<img src="images/화면 기록 2023-09-25 오후 3.17.44.gif"></img>

- 무대인사 크롤링
<img src="images/화면 기록 2023-09-26 오후 12.27.43.gif"></img>
  


</br>

<a name="development-environment"></a>
## 개발 환경

| Service | Dev Language | DevTools |
|----------|----------|----------|
| BE API | Golang | gin web framework, GROM, chromedp, Swagger, JWT, OAuth, PostgreSql |
| Front | javascript | React, Redux |

<a name="cicd-pipeline"></a>
## CI/CD 파이프라인
<img src="images/cloudtype_cicd.PNG"></img>
