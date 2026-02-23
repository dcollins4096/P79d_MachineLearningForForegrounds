#!/bin/csh
set echo
set name = "ms"



set this_latex = pdflatex
if( `grep  "^\\pdftrue" $main |wc -l` == 1 ) set this_latex = pdflatex
if( `grep  "^\\pdffalse" $main |wc -l` == 1 ) set this_latex = latex
if( -e /usr/bin/acroread ) then
 set run = "acroread"
else
 set run = "open"
endif    
rm -f *.log *.out *.aux *.dvi *~ *.bbl *.blg
rm -f *.ps $name.pdf *.bak *.toc *.lof *.lot
set make_only = no
set clean_only = no
while( $#argv )
  switch ( $1 )
    case "o":
      set make_only = "yes"
      shift 
      breaksw
    case "c":
      set clean_only = "yes"
      shift
      breaksw
    default:
      set name = $1
      shift
      breaksw
  endsw
end

if( $clean_only == "yes" ) then
    exit
endif

if( $make_only == "yes" ) then
    $this_latex $name
else
    
    ls -l
    
    $this_latex $name
    if( $status != 0 ) then
	echo "done borked sompin. Exiting."
	exit
    endif
    if ($1 != nb ) then
    bibtex $name
    $this_latex $name
    $this_latex $name
        mv ms.bbl derp.bbl
        #egrep -v "REVTEX41Control|????|apsrev41Control|^08. 1" derp.bbl > ms.bbl
        egrep -v "^ \?\?\?\?|^08. 1" derp.bbl > ms.bbl
    $this_latex $name
    endif
    if( $1 != "n" ) then
	if( -e $name.dvi ) then
	echo "convert to pdf"
	    dvips -Ppdf -t letter -o $name.ps $name.dvi
	    ps2pdf $name.ps $name.pdf

	    if( $status != 0 )then
		print "dvipdf borked.  That's *really* too bad."
		exit
	    endif
	endif
	if( $1 != "n" && ( -e $name.pdf ) ) $run $name.pdf &
    endif	    
endif
# the end
