# Projekat iz predmeta Napredne tehnike programiranja
## _Sistem e-punjača_

**Aplikacija za lakše pronalaženje/rezervaciju električnih punjača za Vaš automobil/bicikl/trotinet.**

## Funkcionalnosti

 - Neregistrovani korisnik
   - Pretraga punjača unošenjem grada (adrese) ili pretragom na mapi
   - Uvid o osnovnim informacijama punjača (cena po kWh, koliko ima slobodnih mesta, kolika je brzina punjenja, ocena, tip priključka punjača)

 - Registrovani korisik
    - _Sve funkcionalnosti neregistrovanog korisnika_
    - Registracija
    - Prijava
    - Pregled/Izmena svojih vozila (električni automobil/bicikl/trotinet)
    - Pregled/Izmena svog profila
    - Mogućnost da sistem na osnovu lokacije korisika odredi najbliži punjač 
    - Rezervacija i otkazivanje punjenja na određenom punjaču
    - Recenzije (komentari) - uz zabranu nekulturnih komentara (email potvrda o zabrani pristupa nalogu korisnicima koji imaju vise od 3 nekulturna komentara) 
    - Sistem nagrađivanja (bodovi) - ukoliko korisnik često koristi određen punjač (grupu punjača) dobija bonuse od distributera
      - BRONZE - 3%; SILVER - 5%; GOLD - 10%; PREMIUM - 15% (popust);
    - Ocenjivanje punjača
    
    _Proširenje za diplomski_
    - Planiranje putovanja (unošenje polazišta - unošenje odredišta, sistem vraća potrošnju energije i optimalnu rutu, kao i estimirano vreme putovanja)
    - Klasifikator nekulturnih komentara (npr. SVM binarni klasifikator)
    

 - Administrator
    - Unošenje novih/Izmena postojećih stanica (e-punjača)
    - Banovanje korisnika - zabrana pristupa korisničkom nalogu - *u slučaju proširenja za diplomski sistem automatski vrši zabranu na osnovu izlaza klasifikatora
    - Izvještaj o korisnicima (banovani - nekulturni komentari, nagrađeni korisnici po kategorijama)
    - Izvještaji o poslovanju (potrošnja energije / prihod, grafikoni)

## Arhitektura 
  - Servisi -  *_moguća izmena pre početka implementacije_
    - Korisnički servis - Go
    - Email servis - Go
    - Rezervacijski servis - Go
    - Servis za izveštaje - Rust
    - Recenzijski servis - Go
    - Traveling-plan servis - Rust
    - Klijentska web aplikacija - Angular (Zamišljeno da bude u Reactu, međutim nisam imao dovoljno vremena)
    
    _Baza_
    - Podaci će biti čuvani u SQL bazi (verovatno PostgreSQL)
