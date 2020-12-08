define(["jquery", "easy-admin", "dragula"], function ($, ea, dragula) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'task/index',
        add_url: 'task/add.html',
        edit_url: 'api/task-edit.json',
        delete_url: 'api/task-delete.json',
        // card_url: 'api/task.json',
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
                    {field: 'Id', width: 80, title: 'ID', sort: true},
                    {field: 'Title', minWidth:150, title: '任务', sort: true},
                    {field: 'Status', width: 80, title: '状态', search: 'select', selectList:{0:"未开始",1:"进行中",2:"已完成"},templet: ea.table.list, sort:true},
                    {field: 'Progress', width: 150, title: '进度',templet: ea.table.percent},
                    {field: 'Owner', width: 120, title: '负责人'},
                    {field: 'Start_time', width: 180, title: '开始时间', sort: true},
                    {field: 'End_time', width: 180, title: '结束时间', sort: true},
                    {field: 'Finish_time', width: 180, title: '完成时间', sort: true},
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
            dragula([$("#todo").get(0), 
                $("#doing").get(0), 
                $("#done").get(0)], {
                    accepts: function (el, target, source, sibling) {
                        if ((source.id == 'todo' && target.id == 'done')
                            || (source.id == 'doing' && target.id == 'todo')
                            || (source.id == 'done' && target.id == 'doing')
                            || (source.id == 'done' && target.id == 'todo')
                            ) {
                            return false;
                        }
                        return true;
                    }
                })
            .on('drag', function (el) {
                // console.log('drag',el)
            }).on('drop', function (el, target, source)  {
                // console.log('drop',el.id, target.id, source.id)
                // todo, update task status

                ea.request.post(
                    {
                        url: '/task/modify/'+el.id,
                        data: {
                            status: target.id
                        }
                    }, function (res) {
                        res.data = res.data || [];
                    }
                );

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