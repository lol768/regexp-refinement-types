FROM raabf/latex-versions:texlive2017
COPY specification /specification
RUN apt-get update
RUN apt-get -y install fonts-font-awesome python-pygments
RUN tlmgr install minted
WORKDIR /specification
CMD ["xelatex", "-shell-escape", "-interaction=nonstopmode", "spec.tex"]
