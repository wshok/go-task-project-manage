define(["jquery", "easy-admin"], function ($, ea) {


    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'api/task.json',
        add_url: 'task/add.html',
        edit_url: 'api/task-edit.json',
        delete_url: 'api/task-delete.json',
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
    };

    return Controller;
});