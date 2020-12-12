'use strict';

/*eslint-disable*/

var ScheduleList = [];

// var SCHEDULE_CATEGORY = [
//     'milestone',
//     'task'
// ];

function ScheduleInfo() {
    this.id = null;
    this.calendarId = null;

    this.title = null;
    this.body = null;
    this.isAllday = true;
    this.start = null;
    this.end = null;
    this.category = 'allday';
    this.dueDateClass = '';

    this.color = null;
    this.bgColor = null;
    this.dragBgColor = null;
    this.borderColor = null;
    this.customStyle = '';

    this.isFocused = false;
    this.isPending = false;
    this.isVisible = true;
    this.isReadOnly = true;
    this.goingDuration = 0;
    this.comingDuration = 0;
    this.recurrenceRule = '';
    this.state = '';

    this.raw = {
        memo: '',
        hasToOrCc: false,
        hasRecurrenceRule: false,
        location: null,
        class: 'public', // or 'private'
        creator: {
            name: '',
            avatar: '',
            company: '',
            email: '',
            phone: ''
        }
    };
}

function generateTime(schedule, renderStart, renderEnd) {
    var startDate = moment(renderStart.getTime())
    var endDate = moment(renderEnd.getTime());
    var diffDate = endDate.diff(startDate, 'days'); // 日期差

    startDate.add(chance.integer({min: 0, max: diffDate}), 'days');
    schedule.start = startDate.toDate();

    endDate = moment(startDate);
    endDate.add(chance.integer({min: 0, max: 3}), 'days');

    schedule.end = endDate.toDate();
}

function generateRandomSchedule(renderStart, renderEnd) {

    var calendar = CalendarList[chance.integer({min: 0, max: 7})]

    var schedule = new ScheduleInfo();

    schedule.id = chance.guid();
    schedule.title = chance.sentence({words: 3});

    generateTime(schedule, renderStart, renderEnd);

    schedule.attendees = ['zhangsan'];
    schedule.state = chance.bool({likelihood: 20}) ? 'Free' : 'Busy';

    schedule.calendarId = calendar.id;
    schedule.color = calendar.color;
    schedule.bgColor = calendar.bgColor;
    schedule.dragBgColor = calendar.dragBgColor;
    schedule.borderColor = calendar.borderColor;

    ScheduleList.push(schedule);
}

function generateSchedule(viewName, renderStart, renderEnd) {
    ScheduleList = [];
    for (var i=0; i < 10; i += 1) {
        generateRandomSchedule(renderStart, renderEnd);
    }
    console.log(ScheduleList)
}
