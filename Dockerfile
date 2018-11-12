FROM raabf/latex-versions:texlive2017
COPY specification /specification
RUN apt-get update
RUN apt-get -y install xzdec fonts-font-awesome python-pygments
RUN tlmgr init-usertree
RUN tlmgr update --all
RUN tlmgr install algorithms avantgar babel babel-french bookman censor charter cm-super cmextra collection-fontsrecommended courier dashrule ec euro euro-ce eurosym exam fpl helvetic invoice labels lineno lm lm-math manfnt-font marvosym mathpazo mflogo-font minted natbib ncntrsbk opensans palatino pxfonts rsfs symbol tex-gyre tex-gyre-math times tipa titling txfonts utopia wasy wasy2-ps wasysym zapfchan zapfding
WORKDIR /specification
CMD ["xelatex", "-shell-escape", "-interaction=nonstopmode", "spec.tex"]
