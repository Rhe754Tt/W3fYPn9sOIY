// 代码生成时间: 2025-10-12 03:27:26
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"

    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()

    // Set the view engine to handle .tmpl files
    app.RegisterView(iris.HTML("./templates", ".tmpl"))

    // Define route for modal dialog
    app.Get("/modal", func(ctx iris.Context) {
        // Render the modal dialog template
        err := ctx.View("modal.tmpl")
        if err != nil {
            // Handle error if template cannot be rendered
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Internal Server Error")
            log.Printf("Error rendering modal template: %v", err)
            return
        }
    })

    // Start the IRIS server
    log.Fatal(app.Listen(":8080"))
}

// This function reads the content of the template file
// and places it into a string. It's used for demonstration purposes.
func loadTemplateContent(filename string) string {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Printf("Error reading template file: %v", err)
        return ""
    }
    return string(content)
}

// Define the template content for modal.tmpl
const modalTemplateContent = `
<!-- templates/modal.tmpl -->
<div id="myModal" class="modal fade" role="dialog">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal">&times;</button>
                <h4 class="modal-title">Modal Header</h4>
            </div>
            <div class="modal-body">
                <p>Some text in the modal.</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
            </div>
        </div>
    </div>
</div>
<script>
// JavaScript for modal dialog
$(document).ready(function(){
    $('#myModal').modal('show');
});
</script>
`
