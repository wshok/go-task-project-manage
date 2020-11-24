define(["jquery", "easy-admin", "dragula"], function ($, ea, dragula) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'api/task.json',
        add_url: 'task/add.html',
        edit_url: 'api/task-edit.json',
        delete_url: 'api/task-delete.json',
        card_url: 'api/task.json',
    };

    var Controller = {
        index: function () {
            ea.table.render({
                init: init,
                // toolbar: [
                //     [{
                //         text: '添加',
                //         url: init.add_url,
                //         method: 'open',
                //         auth: 'add',
                //         class: 'layui-btn layui-btn-normal layui-btn-sm',
                //         icon: 'fa fa-plus ',
                //         extend: 'data-full="true"',
                //     }],
                //     'delete'
                // ],
                cols: [[
                    {type: "checkbox", width: 50},
                    {field: 'id', width: 80, title: 'ID', sort: true},
                    {field: 'title', minWidth:150, title: '任务', sort: true},
                    {field: 'status', width: 80, title: '状态', search: 'select', selectList:{0:"未开始",1:"进行中",2:"已完成"},templet: ea.table.list, sort:true},
                    {field: 'progress', width: 150, title: '进度',templet: ea.table.percent},
                    {field: 'owner', width: 120, title: '负责人'},
                    {field: 'start_time', width: 180, title: '开始时间', sort: true},
                    {field: 'end_time', width: 180, title: '结束时间', sort: true},
                    {field: 'finish_time', width: 180, title: '完成时间', sort: true},
                    {
                        width: 150,
                        title: '操作',
                        templet: ea.table.tool,
                        operat: ['edit', 'delete']
                    }
                ]],
            });

            ea.listen();
        },
        add: function () {
            ea.listen();
        },
        edit: function () {
            ea.listen();
        },
        card: function () {
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
                // todo, update task status
            }).on('over', function (el, container) {
                // console.log('over',el, container)
            }).on('out', function (el, container) {
                // console.log('out',el, container)
            });

            ea.listen();
        },
    };

    return Controller;
});