use aws_sdk_ec2::{Client, Error};
use std::result::Result;

pub async fn list_vpc() -> Result<(), Error> {
    let config = aws_config::from_env().load().await;
    let client = Client::new(&config);

    let resp = client.describe_vpcs().send().await;
    if resp.is_err() {
        println!("error: {:?}", resp);
    }

    for vpc in resp.unwrap().vpcs.unwrap_or_default() {
        println!(
            "id: {}      cidr: {}",
            vpc.vpc_id.unwrap_or_default(),
            vpc.cidr_block.unwrap_or_default()
        )
    }
    println!("hello");
    Ok(())
}
