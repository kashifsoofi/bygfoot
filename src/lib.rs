use std::env;
use std::fs;
use std::path::PathBuf;
use std::path::Path;
extern crate dirs;

pub fn os_is_unix() -> bool {
    env::consts::FAMILY == "unix"
}

fn get_current_dir() -> PathBuf {
    let mut dir = env::current_exe().expect("get current path failed");
    dir.pop();
    dir
}

fn get_bygfoot_dir() -> PathBuf {
    if os_is_unix() {
        let mut home = dirs::home_dir().expect("get home failed");
        home.push(".bygfoot");
        return home;
    }

    get_current_dir()
}

#[no_mangle]
pub extern "C" fn save_hint_number(hint_number: i32) {
    let mut path = get_bygfoot_dir();
    path.push("hint_num");
    
    fs::write(path, &hint_number.to_string()).expect("write failed");
}

#[no_mangle]
pub extern "C" fn load_hint_number() -> i32 {
    let mut path = get_bygfoot_dir();
    path.push("hint_num");

    if !Path::new(&path).exists() {
        return 0;
    }

    let data = fs::read_to_string(path).expect("unable to read file");
    data.parse::<i32>().unwrap()
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[cfg(target_family = "unix")]
    fn test_os_is_unix() {
        let result = os_is_unix();
        assert!(result);
    }

    #[test]
    #[cfg(target_family = "windows")]
    fn test_os_is_unix() {
        let result = os_is_unix();
        assert!(!result);
    }
}
