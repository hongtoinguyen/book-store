db.createUser(
  {
    user: "admin",
    pwd: "bs_1234",
    roles: ["readWrite", "dbAdmin"]
  }
);

db.createCollection("customers");
db.createCollection("books");
db.createCollection("orders");
db.createCollection("payments");
db.createCollection("carts");
db.createCollection("shipping_infos");
db.createCollection("comments");
db.createCollection("staffs");