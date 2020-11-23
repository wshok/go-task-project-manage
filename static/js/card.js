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
            dragula([$("#drag-list-1").get(0), 
                $("#drag-list-2").get(0), 
                $("#drag-list-3").get(0)], {
                    accepts: function (el, target, source, sibling) {
                        if ((source.id == 'drag-list-1' && target.id == 'drag-list-3')
                            || (source.id == 'drag-list-2' && target.id == 'drag-list-1')
                            || (source.id == 'drag-list-3' && target.id == 'drag-list-2')
                            || (source.id == 'drag-list-3' && target.id == 'drag-list-1')
                            ) {
                            return false;
                        }
                        return true;
                    }
                })
            .on('drag', function (el) {
                // console.log('drag',el)
            }).on('drop', function (el, target, source, sibling)  {
                console.log('drop',el, target, source, sibling)
                // todo
            }).on('over', function (el, container) {
                // console.log('over',el, container)
            }).on('out', function (el, container) {
                // console.log('out',el, container)
            });

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