BEGIN TRANSACTION;

INSERT INTO `Item` (ID,Name,Description,Image,Location,Category,Price,Quantity) VALUES 

(1,'Canon EOS 80D DSLR','Eine leistungsstarke, vielseitige und reaktionsschnelle Kamera zum Erkunden der eigenen Kreativität. Die EOS 80D bietet innovative Technologien, die Ihnen in jeder Situation dabei helfen, erstklassige Ergebnisse zu erreichen. Daher überzeugt sie bei Sport- wie bei Porträtaufnahmen, bei der Landschafts-, Street-, Reise- und Low-Light-Fotografie sowie mit hochwertigen Videoaufnahmen.','/img/Canon_EOS_80D.jpg',1,1,42,0),

 (2,'Nikon D7200','Machen Sie den Schritt in die Spitzenklasse mit der D7200. Diese DX-Format-Spiegelreflexkamera ist reaktionsschnell, liefert atemberaubende Fotos, glänzt auch bei Videoaufnahmen und ist vollständig vernetzt. Ihre Top-Ausstattung mit professionellen Features macht sie schnell und vielseitig und lässt sie die Erwartungen jedes Fotografen erfüllen, der mit einem leichten System Außergewöhnliches schaffen will.','/img/Nikon_D7200.jpg',1,1,42,23),

 (3,'Behringer XM8500','Das Behringer XM8500 Ultravoice dynamische Mikrofon wurde entwickelt, um Ihre Sounds mit hervorragender Präzision zu erfassen. Die Nieren-Richtcharakteristik liefert ausgezeichnete Trennung der Schallquellen bei Minimierung von Feedback und Nebengeräuschen. Mit dem Behringer Ultravoice XM8500 Mikrofon und seiner mitgeliefterten Klemme und seiner schlagfesten Aufbewahrungsbox Sind Sie für jeden Auftritt und Studiosession bestens ausgerüstet.','/img/Behringer_XM8500.jpg',2,2,16,1),

 (4,'Alto DVM5','DVM5 Handheld Dynamisches Mikrofon. Das Alto DVM5 ist ein herausragendes Allround-Gesangsmikrofon, ideal für alles, von energetischer Countrymusik über schwülstige, emotionale Jazz-Performances bis hin zu hochkarätigen Rock-Lead-Sängern. Die Hauptmerkmale der DVM5 sind die Flexibilität und die weitreichenden Möglichkeiten. Sie zeichnen sich durch einen sehr niedrigen Grundrauschen, die Fähigkeit, sehr hohe Schallpegel ohne störende Übersteuerung oder Verzerrung auszunutzen, und einen breiten, weichen Frequenzgang aus. Das DVM5 ist extrem tolerant gegenüber wechselnden klimatischen Bedingungen und eignet sich daher besonders gut für den Außeneinsatz. Außerdem ist es mit einem Mikrofonclip zur bequemen Befestigung ausgestattet.','/img/Alto_DVM5.jpg',2,2,16,5),

 (5,'Stairville Outdoor Stage Par','Der Scheinwerfer ist eine Leuchte, in der das durch ein Leuchtmittel (z. B. Glühlampe, Gasentladungslampe, Bogenlampe, Leuchtdiode) erzeugte Licht durch scharfe Bündelung der Lichtstrahlen (Reflexion oder Brechung) in eine Richtung gelenkt wird.','/img/Stairville_Outdoor_Stage_Par.jpg',3,3,20,2),

 (6,'AKG K 271 MK II','Der AKG K271 MKII ist ein geschlossener Studiokopfhörer mit ohrumschließenden Ohrpolstern. Er ist vielseitig einsetzbar und eignet sich für den Einsatz im Studio und bei Live-Produktionen. Bei Studiokopfhörern wird besonders Wert auf eine neutrale Wiedergabe gelegt. Erst so lässt sich im Tonstudio die Qualität einer Aufnahme richtig beurteilen.','/img/AKG_K_271_MK_II.jpg',4,4,150,63),

 (7,'Beyerdynamic Custom Studio','Der CUSTOM Studio ist ein geschlossener, dynamischer Kopfhörer für den professionellen Studioeinsatz. Über den vierstufigen „CUSTOM Studio Sound Slider“ lässt sich das Klangprofil je nach Situation von einem linearen oder analytischen Klang bis zu einem reichhaltigen, kräftigen Bass anpassen.','/img/Beyerdynamic_Custom_Studio.jpg',4,4,50,5),

 (8,'Stairville PAR 64','Der Scheinwerfer ist eine Leuchte, in der das durch ein Leuchtmittel (z. B. Glühlampe, Gasentladungslampe, Bogenlampe, Leuchtdiode) erzeugte Licht durch scharfe Bündelung der Lichtstrahlen (Reflexion oder Brechung) in eine Richtung gelenkt wird.','/img/Stairville_PAR_64.jpg',3,3,20,1),

 (9,'Power Dynamics PDSM8 Studio-Monitore','Die Power Dynamics PDSM8 Studio Monitore überzeugen durch präzise und unverzerrte Klangwiedergabe im Rahmen von Nahfeld- und Midfield-Anwendungen. Hervorragende Frequenztrennung mittels integrierter, elektronischer Weiche und eine effektive Bassreflexbauweise ermöglichen den anspruchsvollen Nutzer die leichte Ortung sämtlicher Instrumente und Nuancen im akustischen Klangraum und deren perfekte Abstimmung aufeinander im Produktionsprozess.','/img/PDSM8_Studio_Monitore.jpg',5,5,25,5),

 (10,'Power Dynamics Galax-5','Der Power Dynamics Galax-5 ist ein aktiver Nahfeld-Lautsprecher mit breitem Frequenzbereich und einem sehr natürlichen, transparenten Klangbild. Um ein besseres Signal-Rausch-Verhältnis zu gewährleisten verfügt er über eine integrierte, elektronische Frequenzweiche für geringere Verzerrungen.','/img/Power_Dynamics Galax_5.jpg',5,5,25,3),

 (11,'Millenium BS-2222 Pro Set','Um den optimalen Klang aus der Anlage zu holen, spielt die richtige Aufstellung der Boxen z.B. mit Boxenständer eine entscheidende Rolle. Egal ob große Basslautsprecher, Front- oder Rearboxen.','/img/Millenium_BS-2222_Pro_Set.jpg',6,6,15,6),

 (12,'Millenium KB-2006 Keyboardbank','Millenium KB-2006 Keyboardbank, große Sitzfläche , Höhenverstellbar von 49 cm bis 64 cm, Gewicht 5,5 kg, Bezug, max. Belastungsgewicht 120 kg, Kunstleder ...','/img/Millenium_KB-2006_Keyboardbank.jpg',6,6,15,2);
COMMIT;
