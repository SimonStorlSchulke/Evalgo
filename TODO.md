## TODO:
- HTML Validator
- Studiengang Attribute (MIB, MKB...)
- (Bilder Upload für Posts)
- Gruppenfarben manuell zuweisbar machen
- PDF Download
- Javascript savety in user posts
- Links zu jeweiliger Aufgabe in Notentabelle
- Config- und Nutzerverwaltungssseite
    - UpdateConfif() nutzen damit Server nicht neugestarted werden muss für Config Update
- Printf für intToString() nutzen

### BUGS:
- Gruppenfarben im geschlossenen Kurs falsch angezeigt (weil Grupenfarbe automatisch den angezeigten Studenten zugewiesen wird und im geschl. Kurs nur einer angezeigt wird)

### Hässlicher Code:
- Farbzuweisung in mainsite.go
- Code "entdenglischen"
- Posts objektorientierter gestalten

- long term - Refactor für besseres error handling und klare Struktur