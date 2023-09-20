(function () {
    'use strict';
    const RECURSION_DURATION = 500;
    let recursion = () => {
        let extraTime = 0;
        try {
            let done = false;
            let video = document.querySelector('#player_html5_api');
            if (video) {
                video.playbackRate = 2;
                video.muted = true;
                if (!video.ended)
                    video.play();
                else
                    video.pause();
                if (video.ended)
                    done = true;
                let quizLayer = document.querySelector('#quizLayer');
                if (quizLayer && quizLayer.style.display != 'none') {
                    if (done) {
                        setTimeout(() => {
                            document.querySelectorAll('.layui-layer-shade').forEach(e => e.style.display = 'none');
                        }, RECURSION_DURATION << 1);
                    };
                    let source = JSON.parse(document.querySelector('div[uooc-video]').getAttribute('source'));
                    let quizList = source.quiz;
                    let quizIndex = 0;
                    let quizQuestion = document.querySelector('.smallTest-view .ti-q-c').innerHTML;
                    for (let i = 0; i < quizList.length; i++) {
                        if (quizList[i].question == quizQuestion) {
                            quizIndex = i;
                            break;
                        };
                    };
                    let quizAnswer = eval(quizList[quizIndex].answer);
                    let quizOptions = quizLayer.querySelector('div.ti-alist');
                    for (let ans of quizAnswer) {
                        let labelIndex = ans.charCodeAt() - 'A'.charCodeAt();
                        quizOptions.children[labelIndex].click();
                    };
                    quizLayer.querySelector('button').click();
                    extraTime = 1000;
                };
                if (!done) {
                    if (video.paused) {
                        video.play();
                    } else {
                        document.querySelectorAll('.layui-layer-shade, #quizLayer').forEach(e => e.style.display = 'none');
                    };
                };
            };
            if (!done) {
                setTimeout(recursion, RECURSION_DURATION + extraTime);
            } else if (video) {
                let current_video = document.querySelector('.basic.active');
                let next_part = current_video.parentNode;
                let next_video = current_video;
                // 定义判断是否视频的函数
                let isVideo = node => Boolean(node.querySelector('span.icon-video'));
                // 定义是否可返回上一级目录的函数
                let canBack = () => {
                    return Boolean(next_part.parentNode.parentNode.tagName === 'LI');
                };
                // 定义更新至后续视频的函数
                let toNextVideo = () => {
                    next_video = next_video.nextElementSibling;
                    while (next_video && !isVideo(next_video)) {
                        next_video = next_video.nextElementSibling;
                    };
                };
                // 定义判断是否存在视频的函数
                let isExistsVideo = () => {
                    let _video = next_part.firstElementChild;
                    while (_video && !isVideo(_video)) {
                        _video = _video.nextElementSibling;
                    };
                    return Boolean(_video && isVideo(_video));
                };
                // 定义判断是否存在后续视频的函数
                let isExistsNextVideo = () => {
                    let _video = current_video.nextElementSibling;
                    while (_video && !isVideo(_video)) {
                        _video = _video.nextElementSibling;
                    };
                    return Boolean(_video && isVideo(_video));
                };
                // 定义检查文件后是否存在后续目录的函数
                let isExistsNextListAfterFile = () => {
                    let part = next_part.nextElementSibling;
                    return Boolean(part && part.childElementCount > 0);
                };
                // 定义更新文件后的后续目录的函数
                let toNextListAfterFile = () => {
                    next_part = next_part.nextElementSibling;
                };
                // 定义返回上一级的函数
                let toOuterList = () => {
                    next_part = next_part.parentNode.parentNode;
                };
                // 定义返回主条目的函数
                let toOuterItem = () => {
                    next_part = next_part.parentNode;
                };
                // 定义检查列表后是否存在后续目录的函数
                let isExistsNextListAfterList = () => {
                    return Boolean(next_part.nextElementSibling);
                };
                // 定义进入列表后的后续目录的函数
                let toNextListAfterList = () => {
                    next_part = next_part.nextElementSibling;
                };
                // 定义展开目录的函数
                let expandList = () => {
                    next_part.firstElementChild.click();
                };
                // 定义进入展开目录的第一个块级元素的函数
                let toExpandListFirstElement = () => {
                    next_part = next_part.firstElementChild.nextElementSibling;
                    if (next_part.classList.contains('unfoldInfo')) {
                        next_part = next_part.nextElementSibling;
                    };
                };
                // 定义判断块级元素是否目录列表的函数
                let isList = () => {
                    return Boolean(next_part.tagName === 'UL');
                };
                // 定义目录列表的第一个目录的函数
                let toInnerList = () => {
                    next_part = next_part.firstElementChild;
                };
                // 定义进入文件列表的第一个视频的函数
                let toFirstVideo = () => {
                    next_video = next_part.firstElementChild;
                    while (next_video && !isVideo(next_video)) {
                        next_video = next_video.nextElementSibling;
                    };
                };
                // 定义模式
                let mode = {
                    FIRST_VIDEO: 'FIRST_VIDEO',
                    NEXT_VIDEO: 'NEXT_VIDEO',
                    LAST_LIST: 'LAST_LIST',
                    NEXT_LIST: 'NEXT_LIST',
                    INNER_LIST: 'INNER_LIST',
                    OUTER_LIST: 'OUTER_LIST',
                    OUTER_ITEM: 'OUTER_ITEM',
                }
                // 定义搜索函数
                let search = (_mode) => {
                    switch (_mode) {
                        case mode.FIRST_VIDEO:
                            if (isExistsVideo()) {
                                toFirstVideo();
                                next_video.click();
                                setTimeout(recursion, RECURSION_DURATION);
                            } else if (isExistsNextListAfterFile()) {
                                search(mode.LAST_LIST);
                            };
                            break;
                        case mode.NEXT_VIDEO:
                            if (isExistsNextVideo()) {
                                toNextVideo();
                                next_video.click();
                                setTimeout(recursion, RECURSION_DURATION);
                            } else if (isExistsNextListAfterFile()) {
                                search(mode.LAST_LIST);
                            } else {
                                search(mode.OUTER_ITEM);
                            };
                            break;
                        case mode.LAST_LIST:
                            toNextListAfterFile();
                            toInnerList();
                            search(mode.INNER_LIST);
                            break;
                        case mode.NEXT_LIST:
                            toNextListAfterList();
                            search(mode.INNER_LIST);
                            break;
                        case mode.INNER_LIST:
                            expandList();
                            (function waitForExpand() {
                                if (next_part.firstElementChild.nextElementSibling) {
                                    toExpandListFirstElement();
                                    if (isList()) {
                                        toInnerList();
                                        search(mode.INNER_LIST);
                                    } else {
                                        search(mode.FIRST_VIDEO);
                                    };
                                } else {
                                    setTimeout(waitForExpand, RECURSION_DURATION);
                                };
                            })();
                            break;
                        case mode.OUTER_LIST:
                            toOuterList();
                            if (isExistsNextListAfterList()) {
                                search(mode.NEXT_LIST);
                            } else if (canBack()) {
                                search(mode.OUTER_LIST);
                            } else {
                                // perhaps there is no next list
                            };
                            break;
                        case mode.OUTER_ITEM:
                            toOuterItem();
                            if (isExistsNextListAfterList()) {
                                toNextListAfterList();
                                search(mode.INNER_LIST);
                            } else if (canBack()) {
                                search(mode.OUTER_LIST);
                            } else {
                                // perhaps there is no list
                            };
                            break;
                        default:
                            break;
                    };
                };
                try {
                    search(mode.NEXT_VIDEO);
                } catch (err) {
                    console.error(err);
                };
            };
        } catch (err) {
            console.error(err);
        };
    };
    recursion();
    console.info('Auto Mooc Init');
})();

function selectClass() {
    let lv1 = document.querySelector('.basic.active');
    let lv1Part = lv1.parentNode;
    let lv1Sub = lv1Part.querySelector('ul');
    let lv1Res = lv1Part.querySelector('.resourcelist');
    if (lv1Res != null) {
        let res = lv1Res.firstElementChild
        res.click()
    } else if (lv1Sub != null) {
        let lv2Part = lv1Sub.firstElementChild
        let lv2 = lv2Part.querySelector('.basic');
        lv2.click()
        setTimeout(() => {
            let lv2Sub = lv2Part.querySelector('ul');
            let lv2Res = lv2Part.querySelector('.resourcelist');
            if (lv2Res != null) {
                let res = lv2Res.firstElementChild
                res.click()
            } else if (lv2Sub != null) {
                let lv3Part = lv2Sub.firstElementChild
                let lv3 = lv3Part.querySelector('.basic');
                lv3.click()
            }
        }, 500)
    }
}