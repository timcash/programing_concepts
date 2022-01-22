// a program to read in a list of community names and their members from a JSON file

// a struct to hold the data
#[derive(Debug)]
struct Community {
    name: String,
    members: Vec<String>,
}

fn main() {
    // read in the data from the JSON file
    let data = fs::read_to_string("data/communities.json").expect("Failed to read data");

    // parse the JSON data into a vector of structs
    let communities: Vec<Community> = serde_json::from_str(&data).expect("Failed to parse data");

    // print the data
    for community in &communities {
        println!("Community: {}", community.name);
        println!("Members:");
        for member in &community.members {
            println!("\t{}", member);
        }
    }
}