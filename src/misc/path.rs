use std::ffi::CStr;
use xplm_sys::XPLMGetSystemPath;
pub fn get_system_path() -> String {
    // Define a buffer size for the path (X-Plane suggests 512 bytes as the maximum size).
    const PATH_MAX_LEN: usize = 512;

    // Allocate a buffer for the path.
    let mut buffer = vec![0u8; PATH_MAX_LEN];

    // Call the C function, passing the buffer as a mutable pointer.
    unsafe {
        XPLMGetSystemPath(buffer.as_mut_ptr() as *mut ::std::os::raw::c_char);
    }

    // Convert the C string in the buffer to a Rust String.
    let c_str = unsafe { CStr::from_ptr(buffer.as_ptr() as *const ::std::os::raw::c_char) };

    // Return the string or handle errors if needed.
    c_str.to_string_lossy().into_owned()
}
