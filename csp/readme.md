

# Csp1978 起源

- Tony Hoare 提出 CSP 的时代背景是什么？
- CSP 1978 理论到底有哪些值得我们研究的地方？
- CSP 1978 理论是否真的就是我们目前熟知的基于通道的同步方式？
- CSP 1978 理论的早期设计存在什么样的缺陷？


# 资料

- Reading CSP 理解顺序进程间通信（Communicating Sequential Processes）
  - https://www.youtube.com/watch?v=Z8ZpWVuEx8c
  - [PPT](https://docs.google.com/presentation/d/1N5skL6vR9Wxk-I82AYs3dlsOsJkUAGJCsb5NGpXWpqo/edit#slide=id.g6571c725ac_0_127)
  - [Ou 2019a] [CSP1978 的 Go 语言实现](https://github.com/changkun/pkg/blob/master/csp/csp.go)
  - ppt： https://docs.google.com/presentation/d/1N5skL6vR9Wxk-I82AYs3dlsOsJkUAGJCsb5NGpXWpqo/edit#slide=id.g6571c725ac_0_127
  - [Hoare 1978] [Hoare, C. A. R. (1978). Communicating sequential processes. Communications of the ACM, 21(8), 666–677.](https://spinroot.com/courses/summer/Papers/hoare_1978.pdf)
  - [Ou 2019b] [第 56 期 channel & select 源码分析](https://github.com/developer-learning/night-reading-go/issues/450)
  - [Ou 2019c] [第 59 期 Real-world Go Concurrency Bugs](https://github.com/developer-learning/night-reading-go/issues/464)




并行的方式：
-  共享内存
- 条件临界区
- 信号量 



如何实现图灵完备性