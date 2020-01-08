package event

func KripaluJiMaharajEvents() []*EventDetail {
	var allEvents []*EventDetail
	kripaluJiJagadGurutam := &EventDetail{
		Day:   14,
		Month: 1,
		Year:  1957,
		Title: "Kripalu Ji Maharaj - Jagadgurutam title day",
		Info: `    He was formally installed as the fifth Jagadguru (world teacher).
   He was 34 years old when given the title on 14 January 1957 by the Kashi Vidvat Parishat, a group of Hindu scholars.
   The Kashi Vidvat Parishat conferred on him the titles Bhaktiyog-Ras-Avtar and Jagadguruttama.
   Followers claim that he is the "fifth original Jagadguru" in the series of Jagadgurus after 
   Śrīpāda Śaṅkarācārya (A.D. 788-820),
   Śrīpāda Rāmānujācārya (1017-1137),
   Śrī Nimbārkācārya and, 
   Śrīpāda Madhvācārya (1239-1319)

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
`,
		URL: "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
	}
	allEvents = append(allEvents, kripaluJiJagadGurutam)

	kripaluJiAppearance := &EventDetail{
		Day:   5,
		Month: 10,
		Year:  1922,
		Title: "Kripalu Ji Maharaj - appearance day",
		Info: `    Kripalu (Sanskrit: जगद्गुरु श्री कृपालु जी महाराज, IAST: Kṛpālu) (5 October 1922 – 15 November 2013) 
    He was a Hindu spiritual leader  from Allahabad (Prayag) - Mangarh, Pratapgarh, India.
    He was the founder of Jagadguru Kripalu Parishat (JKP), 
    a worldwide Hindu non-profit organization with five main ashrams; four in India and one in the United States.
    JKP Radha Madhav Dham is one of the largest Hindu Temple complexes in the Western Hemisphere, and the largest in North America.

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
     `,
		URL: "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
	}
	allEvents = append(allEvents, kripaluJiAppearance)

	kripaluJiDisAppearance := EventDetail{
		Day:   15,
		Month: 11,
		Year:  2013,
		Title: "Kripalu Ji Maharaj - disappearance day",
		Info: `He disappeared 15 November 2013 (aged 91) New Delhi, Delhi, the dense forest near Sharbhang ashram in Chitrakoot.

   He was formally installed as the fifth Jagadguru (world teacher).
   He was 34 years old when given the title on 14 January 1957 by the Kashi Vidvat Parishat, a group of Hindu scholars.
   The Kashi Vidvat Parishat conferred on him the titles Bhaktiyog-Ras-Avtar and Jagadguruttama.
   Followers claim that he is the "fifth original Jagadguru" in the series of Jagadgurus after 
   Śrīpāda Śaṅkarācārya (A.D. 788-820),
   Śrīpāda Rāmānujācārya (1017-1137),
   Śrī Nimbārkācārya and, 
   Śrīpāda Madhvācārya (1239-1319)

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
    `,
		URL: "https://en.wikipedia.org/wiki/Kripalu_Maharaj;https://www.jkyog.org/;http://jkp.org.in/life-story-jagadguruttam/",
	}
	allEvents = append(allEvents, &kripaluJiDisAppearance)
	return allEvents
}