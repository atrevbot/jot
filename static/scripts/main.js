// function handleBeforeUnload(event) {
//     // Cancel the event as stated by the standard.
//     event.preventDefault();
//     // Chrome requires returnValue to be set.
//     event.returnValue = '';
// }
// 
// window.addEventListener("beforeunload", handleBeforeUnload);

if (window.location.pathname === "/close") {
    window.close()
}
