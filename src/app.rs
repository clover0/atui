use crate::input::InputMode;

// App hold state
pub struct App {
    pub input: String,
    pub input_mode: InputMode,
    pub messages: Vec<String>,
}

impl Default for App {
    fn default() -> App {
        App {
            input: String::new(),
            input_mode: InputMode::Normal,
            messages: Vec::new(),
        }
    }
}
