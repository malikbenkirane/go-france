go-france
=

Fake french first and last names.

Raw data from the INSEE:
- [noms (insee 1891-2000)](
  https://www.insee.fr/fr/statistiques/3536630
); archives at `public/insee_noms_2008.txt.zip`
- [prenoms (insee 1900-2022)](
  https://www.insee.fr/fr/statistiques/7633685?sommaire=7635552
); archives at `public/insee_nat_2022.csv.zip` 


Use
-

    go get github.com/malikbenkirane/go-france

Read th docs at [pkg.go.dev/github.com/malikbenkirane/go-france](
  https://pkg.go.dev/github.com/malikbenkirane/go-france
). You will find there an example with `NewSet` on how to load provided
firstnames and lastnames sets concurrently.

Of course this is not perfect and every suggestion is welcome 🙂

Feel free to file issues and pull requests 🚀


Roadmap
-

Open to suggestions... 😊


nat.ipynb
-

Known to work:
extract files from `/public`, and open nat.ipynb in a jupyter lab,
this would generate the files you can find in `/data`.
