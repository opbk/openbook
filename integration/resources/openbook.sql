
INSERT INTO categories (parent_id, path, name, books) VALUES (0, '', 'Бизнес и деньги', 5);
INSERT INTO categories (parent_id, path, name, books) VALUES (1, '1', 'Индустрия', 5);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Агрокультура', 1);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Компьютеры и технологии', 2);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Энергия', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Мода и текстиль', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Туризм и путешествия', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Производство', 1);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Медиа', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Музеи', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Биотехнологии', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Рестораны и еда', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Ретейлинг', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Сервис', 0);
INSERT INTO categories (parent_id, path, name, books) VALUES (2, '1>2', 'Спорт и развлечения', 0);

INSERT INTO publishers (name, description, books) VALUES ('Knopf', '', 1);
INSERT INTO publishers (name, description, books) VALUES ('Simon & Schuster', '', 1);
INSERT INTO publishers (name, description, books) VALUES ('W. W. Norton & Company', '', 1);
INSERT INTO publishers (name, description, books) VALUES ('Grand Central Publishing', '', 1);
INSERT INTO publishers (name, description, books) VALUES ('Back Bay Books', '', 1);

INSERT INTO authors (name, description, books) VALUES ('Sven Beckert', '', 1);
INSERT INTO authors (name, description, books) VALUES ('Walter Isaacson', '', 1);
INSERT INTO authors (name, description, books) VALUES ('Michael Lewis', '', 1);
INSERT INTO authors (name, description, books) VALUES ('Eric Schmidt', '', 1);
INSERT INTO authors (name, description, books) VALUES ('Jonathan Rosenberg', '', 1);
INSERT INTO authors (name, description, books) VALUES ('Beth Macy', '', 1);

INSERT INTO author_books (book_id, author_id) VALUES (1, 1);
INSERT INTO author_books (book_id, author_id) VALUES (2, 2);
INSERT INTO author_books (book_id, author_id) VALUES (3, 3);
INSERT INTO author_books (book_id, author_id) VALUES (4, 4);
INSERT INTO author_books (book_id, author_id) VALUES (4, 5);
INSERT INTO author_books (book_id, author_id) VALUES (5, 6);

INSERT INTO book_categories (book_id, category_id) VALUES (1, 1);
INSERT INTO book_categories (book_id, category_id) VALUES (1, 2);
INSERT INTO book_categories (book_id, category_id) VALUES (1, 3);

INSERT INTO book_categories (book_id, category_id) VALUES (2, 1);
INSERT INTO book_categories (book_id, category_id) VALUES (2, 2);
INSERT INTO book_categories (book_id, category_id) VALUES (2, 4);

INSERT INTO book_categories (book_id, category_id) VALUES (3, 1);
INSERT INTO book_categories (book_id, category_id) VALUES (3, 2);

INSERT INTO book_categories (book_id, category_id) VALUES (4, 1);
INSERT INTO book_categories (book_id, category_id) VALUES (4, 2);
INSERT INTO book_categories (book_id, category_id) VALUES (4, 4);

INSERT INTO book_categories (book_id, category_id) VALUES (5, 1);
INSERT INTO book_categories (book_id, category_id) VALUES (5, 2);
INSERT INTO book_categories (book_id, category_id) VALUES (5, 8);

INSERT INTO prices (type, name) VALUES ('retail', 'Розничная цена');
INSERT INTO book_prices(book_id, price_type_id, price) VALUES (1, 1, 590);
INSERT INTO book_prices(book_id, price_type_id, price) VALUES (2, 1, 995);
INSERT INTO book_prices(book_id, price_type_id, price) VALUES (3, 1, 1295);
INSERT INTO book_prices(book_id, price_type_id, price) VALUES (4, 1, 769);
INSERT INTO book_prices(book_id, price_type_id, price) VALUES (5, 1, 809);

INSERT INTO books (series_id, publisher_id, title, pages, language, release, created, short, description, service_review, critics_review)
	VALUES (0, 1, 'Empire of Cotton: A Global History', 640, 'Английский', '2014-12-02', NOW(), '<p>The epic story of the rise and fall of the empire of cotton, its centrality to the world economy, and its making and remaking of global capitalism.</p><p>Cotton is so ubiquitous as to be almost invisible, yet understanding its history is key to understanding the origins of modern capitalism. Sven Beckert’s rich, fascinating book tells the story of how, in a remarkably brief period, European entrepreneurs and powerful statesmen recast the world’s most significant manufacturing industry, combining imperial expansion and slave labor with new machines and wage workers to change the world.</p>', '<p>The epic story of the rise and fall of the empire of cotton, its centrality to the world economy, and its making and remaking of global capitalism.</p><p>Cotton is so ubiquitous as to be almost invisible, yet understanding its history is key to understanding the origins of modern capitalism. Sven Beckert’s rich, fascinating book tells the story of how, in a remarkably brief period, European entrepreneurs and powerful statesmen recast the world’s most significant manufacturing industry, combining imperial expansion and slave labor with new machines and wage workers to change the world. Here is the story of how, beginning well before the advent of machine production in the 1780s, these men captured ancient trades and skills in Asia, and combined them with the expropriation of lands in the Americas and the enslavement of African workers to crucially reshape the disparate realms of cotton that had existed for millennia, and how industrial capitalism gave birth to an empire, and how this force transformed the world.</p><p>The empire of cotton was, from the beginning, a fulcrum of constant global struggle between slaves and planters, merchants and statesmen, workers and factory owners. Beckert makes clear how these forces ushered in the world of modern capitalism, including the vast wealth and disturbing inequalities that are with us today. The result is a book as unsettling as it is enlightening: a book that brilliantly weaves together the story of cotton with how the present global world came to exist.</p>', '<p>How important is cotton? For starters, there’s a good chance that you’re wearing it right now. That’s true no matter where you live in the world. Cotton is everywhere, has been for a long time, and was the dominant commodity during the early years of our country. It fostered “war capitalism” among European nations. It helped launch the industrial revolution in England. It drove slavery. The story of cotton is the story of modern capitalism, and in Empire of Cotton, author Sven Beckert shows how a worldwide crop that came in multiple forms and was cultivated and produced in many different ways came to be dominated by the late coming Europeans, and later Americans, often through violent means, reshaping both the world economy and the world itself—for better or worse—along the way.</p>', '');

INSERT INTO books (series_id, publisher_id, title, pages, language, release, created, short, description, service_review, critics_review)
VALUES (0, 2, 'The Innovators: How a Group of Hackers, Geniuses, and Geeks Created the Digital Revolution', 560, 'Английский', '2014-11-07', NOW(), '<p>Following his blockbuster biography of Steve Jobs, The Innovators is Walter Isaacson’s revealing story of the people who created the computer and the Internet. It is destined to be the standard history of the digital revolution and an indispensable guide to how innovation really happens.</p><p>What were the talents that allowed certain inventors and entrepreneurs to turn their visionary ideas into disruptive realities? What led to their creative leaps? Why did some succeed and others fail?</p>', '<p>Following his blockbuster biography of Steve Jobs, The Innovators is Walter Isaacson’s revealing story of the people who created the computer and the Internet. It is destined to be the standard history of the digital revolution and an indispensable guide to how innovation really happens.</p><p>What were the talents that allowed certain inventors and entrepreneurs to turn their visionary ideas into disruptive realities? What led to their creative leaps? Why did some succeed and others fail?</p><p>In his masterly saga, Isaacson begins with Ada Lovelace, Lord Byron’s daughter, who pioneered computer programming in the 1840s. He explores the fascinating personalities that created our current digital revolution, such as Vannevar Bush, Alan Turing, John von Neumann, J.C.R. Licklider, Doug Engelbart, Robert Noyce, Bill Gates, Steve Wozniak, Steve Jobs, Tim Berners-Lee, and Larry Page.</p></p>This is the story of how their minds worked and what made them so inventive. It’s also a narrative of how their ability to collaborate and master the art of teamwork made them even more creative.</p><p>For an era that seeks to foster innovation, creativity, and teamwork, The Innovators shows how they happen.</p>', '<p>Many books have been written about Silicon Valley and the collection of geniuses, eccentrics, and mavericks who launched the “Digital Revolution”; Robert X. Cringely Accidental Empires and Michael A. Hiltzik Dealers of Lightning are just two excellent accounts of the unprecedented explosion of tech entrepreneurs and their game-changing success. But Walter Isaacson goes them one better: The Innovators, his follow-up to the massive (in both sales and size) Steve Jobs, is probably the widest-ranging and most comprehensive narrative of them all. Dont let the scope or page-count deter you: while Isaacson builds the story from the 19th century--innovator by innovator, just as the players themselves stood atop the achievements of their predecessors--his discipline and era-based structure allows readers to dip in and out of digital history, from Charles Babbage Difference Engine, to Alan Turing and the codebreakers of Bletchley Park, to Tim Berners-Lee and the birth of the World Wide Web (with contextual nods to influential counterculture weirdos along the way). Isaacson presentation is both brisk and illuminating; while it doesnt supersede previous histories, The Innovators might be the definitive overview, and it certainly one hell of a read.</p>', '');

INSERT INTO books (series_id, publisher_id, title, pages, language, release, created, short, description, service_review, critics_review)
	VALUES (0, 3, 'Flash Boys', 288, 'Английский', '2014-03-31', NOW(), '<p>Flash Boys is about a small group of Wall Street guys who figure out that the U.S. stock market has been rigged for the benefit of insiders and that, post–financial crisis, the markets have become not more free but less, and more controlled by the big Wall Street banks. Working at different firms, they come to this realization separately; but after they discover one another, the flash boys band together and set out to reform the financial markets. This they do by creating an exchange in which high-frequency trading—source of the most intractable problems—will have no advantage whatsoever.</p>', '<p>Flash Boys is about a small group of Wall Street guys who figure out that the U.S. stock market has been rigged for the benefit of insiders and that, post–financial crisis, the markets have become not more free but less, and more controlled by the big Wall Street banks. Working at different firms, they come to this realization separately; but after they discover one another, the flash boys band together and set out to reform the financial markets. This they do by creating an exchange in which high-frequency trading—source of the most intractable problems—will have no advantage whatsoever.</p><p>The characters in Flash Boys are fabulous, each completely different from what you think of when you think “Wall Street guy.” Several have walked away from jobs in the financial sector that paid them millions of dollars a year. From their new vantage point they investigate the big banks, the world’s stock exchanges, and high-frequency trading firms as they have never been investigated, and expose the many strange new ways that Wall Street generates profits.</p><p>The light that Lewis shines into the darkest corners of the financial world may not be good for your blood pressure, because if you have any contact with the market, even a retirement account, this story is happening to you. But in the end, Flash Boys is an uplifting read. Here are people who have somehow preserved a moral sense in an environment where you don’t get paid for that; they have perceived an institutionalized injustice and are willing to go to war to fix it.</p>', '', '');

INSERT INTO books (series_id, publisher_id, title, pages, language, release, created, short, description, service_review, critics_review)
	VALUES (0, 4, 'How Google Works', 304, 'Английский', '2014-09-23', NOW(), '<p>Google Executive Chairman and ex-CEO Eric Schmidt and former SVP of Products Jonathan Rosenberg came to Google over a decade ago as proven technology executives. At the time, the company was already well-known for doing things differently, reflecting the visionary--and frequently contrarian--principles of founders Larry Page and Sergey Brin. If Eric and Jonathan were going to succeed, they realized they would have to relearn everything they thought they knew about management and business.</p>', '<p>Google Executive Chairman and ex-CEO Eric Schmidt and former SVP of Products Jonathan Rosenberg came to Google over a decade ago as proven technology executives. At the time, the company was already well-known for doing things differently, reflecting the visionary--and frequently contrarian--principles of founders Larry Page and Sergey Brin. If Eric and Jonathan were going to succeed, they realized they would have to relearn everything they thought they knew about management and business.</p><p>Today, Google is a global icon that regularly pushes the boundaries of innovation in a variety of fields. HOW GOOGLE WORKS is an entertaining, page-turning primer containing lessons that Eric and Jonathan learned as they helped build the company. The authors explain how technology has shifted the balance of power from companies to consumers, and that the only way to succeed in this ever-changing landscape is to create superior products and attract a new breed of multifaceted employees whom Eric and Jonathan dub "smart creatives." Covering topics including corporate culture, strategy, talent, decision-making, communication, innovation, and dealing with disruption, the authors illustrate management maxims ("Consensus requires dissension," "Exile knaves but fight for divas," "Think 10X, not 10%") with numerous insider anecdotes from Google history, many of which are shared here for the first time.</p><p>In an era when everything is speeding up, the best way for businesses to succeed is to attract smart-creative people and give them an environment where they can thrive at scale. HOW GOOGLE WORKS explains how to do just that.</p>', '<p>An informative and creatively multilayered Google guidebook from the businessman perspective.</p>', '');

INSERT INTO books (series_id, publisher_id, title, pages, language, release, created, short, description, service_review, critics_review)
VALUES (0, 5, 'Factory Man: How One Furniture Maker Battled Offshoring, Stayed Local - and Helped Save an American Town',  464, 'Английский', '2014-07-15', NOW(), '<p>The Bassett Furniture Company was once the world biggest wood furniture manufacturer. Run by the same powerful Virginia family for generations, it was also the center of life in Bassett, Virginia. But beginning in the 1980s, the first waves of Asian competition hit, and ultimately Bassett was forced to send its production overseas.</p>', '<p>The Bassett Furniture Company was once the world biggest wood furniture manufacturer. Run by the same powerful Virginia family for generations, it was also the center of life in Bassett, Virginia. But beginning in the 1980s, the first waves of Asian competition hit, and ultimately Bassett was forced to send its production overseas.</p><p>One man fought back: John Bassett III, a shrewd and determined third-generation factory man, now chairman of Vaughan-Bassett Furniture Co, which employs more than 700 Virginians and has sales of more than $90 million. In FACTORY MAN, Beth Macy brings to life Bassett deeply personal furniture and family story, along with a host of characters from an industry that was as cutthroat as it was colorful. As she shows how he uses legal maneuvers, factory efficiencies, and sheer grit and cunning to save hundreds of jobs, she also reveals the truth about modern industry in America.</p>', '', '');

INSERT INTO users (email, password, name, created, last_enter) VALUES ('netw00rk@gmail.com', '12345', 'UserName', NOW(), NOW());

ALTER SEQUENCE auto_id_orders RESTART WITH 471

INSERT INTO orders (user_id, address_id, status, comment) VALUES (1, 0, 'returned', '');
INSERT INTO book_orders (order_id, book_id) VALUES (1, 1);

INSERT INTO orders (user_id, address_id, status, comment) VALUES (1, 0, 'onhand', '');
INSERT INTO book_orders (order_id, book_id) VALUES (2, 2);

INSERT INTO book_orders (order_id, book_id) VALUES (3, 3);
INSERT INTO orders (user_id, address_id, status, comment) VALUES (1, 0, 'delivered', '');

INSERT INTO book_orders (order_id, book_id) VALUES (4, 4);
INSERT INTO orders (user_id, address_id, status, comment) VALUES (1, 0, 'inprogress', '');

INSERT INTO orders (user_id, address_id, status, comment) VALUES (1, 0, 'new', '');
INSERT INTO book_orders (order_id, book_id) VALUES (5, 5);