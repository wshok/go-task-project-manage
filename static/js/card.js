define(["jquery", "easy-admin", "dragula"], function ($, ea, dragula) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'api/task.json',
        // add_url: 'card/add.html',
        // edit_url: 'api/task-edit.json',
        // delete_url: 'api/task-delete.json',
    };

    var Controller = {
        index: function () {
            dragula([$("#drag-list-1").get(0), $("#drag-list-2").get(0), $("#drag-list-3").get(0)]);

            ea.listen();
        },
        add: function () {
            ea.listen();
        },
        edit: function () {
            ea.listen();
        },
    };

    return Controller;
});