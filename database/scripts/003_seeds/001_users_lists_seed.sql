WITH "new_users" AS
(
  INSERT INTO "user"
              ("role_id", "first_name", "middle_name", "last_name",   "surname",   "picture_url",                                     "email",                    "password")
       VALUES (1,         'Jeremy',     'Alexander',   'Fonseca',     'Blanco',    'http://dummyimage.com/235x100.png/5fa2dd/ffffff', 'f@mail.com',               '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Harmonie',   'Merla',       'Saberton',    'Shoulders', 'http://dummyimage.com/163x100.png/cc0000/ffffff', 'mshoulders1@shop-pro.jp',  '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Roscoe',     'Merry',       'Sibyllina',   'Dixson',    'http://dummyimage.com/145x100.png/cc0000/ffffff', 'mdixson2@typepad.com',     '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Nedda',      'Kristin',     'Lewin',       'Crispin',   'http://dummyimage.com/207x100.png/5fa2dd/ffffff', 'kcrispin3@alexa.com',      '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Noelyn',     'Muriel',      'De Few',      'Fewkes',    'http://dummyimage.com/165x100.png/5fa2dd/ffffff', 'mfewkes4@photobucket.com', '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Austin',     'Skyler',      'Kitchenside', 'Masson',    'http://dummyimage.com/153x100.png/dddddd/000000', 'smasson5@blogspot.com',    '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Shirlene',   'Illa',        'Staynes',     'MacAless',  'http://dummyimage.com/176x100.png/cc0000/ffffff', 'imacaless6@ameblo.jp',     '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Sherrie',    'Hamnet',      'Prestedge',   'Fackney',   'http://dummyimage.com/214x100.png/5fa2dd/ffffff', 'hfackney7@patch.com',      '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Jared',      'Catlaina',    'McFarlane',   'Craighill', 'http://dummyimage.com/212x100.png/ff4444/ffffff', 'ccraighill8@blogs.com',    '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC'),
              (2,         'Eba',        'Raynard',     'Yakovl',      'Gurnett',   'http://dummyimage.com/182x100.png/dddddd/000000', 'rgurnett9@guardian.co.uk', '$2a$10$ODkl/7qJaSbpg025Ddhu5ewRDsKUOk8L2M6YYnzIW4N9t8mAyGRHC')
    RETURNING "user_id", "first_name", "last_name"
)
INSERT INTO "list" ("owner_id", "name", "description")
     SELECT "user_id",
            'today',
            "first_name" || ' ' || "last_name" || '''s today list'
       FROM "new_users"
  UNION ALL
     SELECT "user_id",
            'tomorrow',
            "first_name" || ' ' || "last_name" || '''s tomorrow list'
       FROM "new_users";
