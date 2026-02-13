

files =  introduction.tex main.tex  title.tex



main.pdf: $(files)
	pdflatex main
	bibtex main
	pdflatex main
	pdflatex main

o: $(files)
	pdflatex main

clean:
	-@rm -f *.{aux,toc,dvi,lof,lot,log,lom,bbl,bcf,blg,pdf,ps,out,run.xml} *~

