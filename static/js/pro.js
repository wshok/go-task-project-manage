define(["jquery", "easy-admin"], function ($, ea) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'pro/index',
        add_url: 'pro/add',
        edit_url: 'pro/edit',
        delete_url: 'pro/delete',
    };

    var Controller = {
        index: function () {
            ea.table.render({
                init: init,
                cols: [[
                    {field: 'id', width: 80, title: 'ID', sort: true},
                    {field: 'title', minWidth:300, title: '标题', templet:function(d){
                        return '<a href="/doc/view/'+d.id+'">'+ d.title +'</a>'
                      }},
                    {field: 'remark', width: 120, title: '备注', search: false},
                    {field: 'start_time', width: 120, title: '开始', search: false},
                    {field: 'end_time', width: 120, title: '结束', search: false},
                    {field: 'create_at', width: 180, title: '创建时间', sort: true, search:false},
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