awk -F"\t" '{
row=int((NR-1)/2);
        for(i=1; i <= NF; i++) {
                gsub("[0-9]", "", $i);
                if ($i != "") {
                        gsub(" ", "", $i);
                        gadi[row][i-1] = gadi[row][i-1]""$i;
                }
        }
} END {
        for (i in gadi) {
                for (j in gadi[i]) {
                        if (gadi[i][j]=="") {
                                printf "#|";
                        } else {
                                printf "%s|", gadi[i][j]; 
                        }
                }
                print "";
        }
}' $@

