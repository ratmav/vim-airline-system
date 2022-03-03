extern crate sys_info;

use chrono;
use sys_info::*;

fn main() {
    let load = loadavg().unwrap();

    let mem = mem_info().unwrap();
    let mem_used = mem.total - mem.free;
    let mem_percent = (100 * mem_used)/mem.total;

    let utc_pretty = chrono::offset::Utc::now()
        .format("%b %d %H:%M utc")
        .to_string()
        .to_lowercase();

    println!(
        "[{:.2}load|{}%ram] {}",
        load.one,
        mem_percent,
        utc_pretty
    );
}
