awk -f extract-dumps.awk < $1 | csplit - '/Full thread dump/' '{*}' -z -f 'threaddump' -b '%03d'
