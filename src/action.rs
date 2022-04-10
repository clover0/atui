use crate::aws::resources::{Resources, RESOURCEMAP};
// use crate::aws::resources::resourceMap;
use crate::aws::ec2::vpc;
use crate::aws::ec2::vpc::list_vpc;

pub async fn handler(input: String) -> Result<(), String> {
    let i = RESOURCEMAP.get(&*input);
    let mut res = i.unwrap_or(&Resources::Unknown);
    println!("{:?}", res);

    if list_vpc().await.is_err() {
        return Err("error".to_string());
    };
    Ok(())
}
