define(["jquery", "easy-admin", "dragula"], function ($, ea, dragula) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'task/index',
        add_url: 'task/add',
        edit_url: 'task/edit',
        delete_url: 'task/delete',
    };

    var Controller = {
        index: function () {
            ea.table.render({
                init: init,
                cols: [[
                    {type: "checkbox", width: 50},
                    {field: 'id', width: 80, title: 'ID', sort: true},
                    {field: 'title', minWidth:150, title: '任务', sort: true},
                    {field: 'status', width: 80, title: '状态', search: 'select', selectList:{'todo':"未开始",'doing':"进行中",'done':"已完成"},templet: ea.table.list, sort:true},
                    {field: 'progress', width: 150, title: '进度',templet: ea.table.percent},
                    {field: 'username', width: 120, title: '负责人'},
                    {field: 'start_time', width: 180, title: '开始时间', sort: true, templet:ea.table.date},
                    {field: 'end_time', width: 180, title: '结束时间', sort: true, templet:ea.table.date},
                    {field: 'finish_time', width: 180, title: '完成时间', sort: true,templet:ea.table.date},
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