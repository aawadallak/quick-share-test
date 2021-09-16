db.createUser(
    {
        user: "root",
        pwd: "123456",
        roles: [
            {
                role: "readWrite",
                db: "easyshare"
            }
        ]
    }
);

db.createCollection(
    "shareable",
)