use lazy_static::lazy_static;
use std::borrow::Borrow;
use std::collections::HashMap;

#[derive(Debug)]
pub enum Resources {
    EC2,
    VPC,
    Subnet,
    SecurityGroup,
    Unknown,
}

lazy_static! {
    pub static ref RESOURCEMAP: HashMap<&'static str, Resources> = {
        let mut m = HashMap::new();
        m.insert("ec2", Resources::EC2);
        m.insert("vpc", Resources::VPC);
        m
    };
}
// lazy_static! {
//     static ref HASHMAP: HashMap<u32, &'static str> = {
//         let mut m = HashMap::new();
//         m.insert(0, "foo");
//         m.insert(1, "bar");
//         m.insert(2, "baz");
//         m
//     };
//     static ref COUNT: usize = HASHMAP.len();
//     static ref NUMBER: u32 = times_two(21);
// }
