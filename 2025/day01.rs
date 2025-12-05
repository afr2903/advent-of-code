use std::fs;

fn main() {
    let mut current_position = 50;
    let mut dial_at_zero = 0;
    let puzzle = fs::read_to_string("day01.txt")
        .expect("wtf?");
    for rotation in puzzle.lines() {
        let (direction, distance_str) = rotation.split_at(1);
        let distance: i32 = distance_str.parse().expect("Not valid number");
        // println!("rot {direction}, dis {distance}");
        current_position += if direction == "L" {-distance} else {distance}; 
        current_position = current_position % 100;
        if current_position == 0 {
            dial_at_zero += 1;
        }
        // println!("curr {current_position}");
    }
    println!("Answer: {dial_at_zero}")
    // println!("With text:\n{puzzle}");
}
